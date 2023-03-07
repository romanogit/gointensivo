package database

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/romanogit/gointensivo/internal/entity"
)

type OrderRepositorySuite struct {
	suite.Suite
	DB        *sql.DB
	OrderRepo *OrderRepository
}

func (suite *OrderRepositorySuite) SetupTest() {
	// create database
	var err error
	suite.DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Printf("failed to create database: %v\n", err)
		os.Exit(1)
	}

	// create orders table
	_, err = suite.DB.Exec("CREATE TABLE Orders (id INTEGER PRIMARY KEY, price INTEGER, tax INTEGER, final_price INTEGER)")
	if err != nil {
		fmt.Printf("failed to create table: %v\n", err)
		os.Exit(1)
	}

	suite.OrderRepo = NewOrderRepository(suite.DB)
}

func (suite *OrderRepositorySuite) TearDownTest() {
	// remove orders table
	_, err := suite.DB.Exec("DROP TABLE IF EXISTS Orders")
	if err != nil {
		fmt.Printf("failed to drop table: %v\n", err)
		os.Exit(1)
	}

	suite.DB.Close()
}

func (suite *OrderRepositorySuite) TestOrderRepository_Save() {
	order := &entity.Order{
		ID:         "1",
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}

	err := suite.OrderRepo.Save(order)
	assert.NoError(suite.T(), err)

	var count int
	err = suite.DB.QueryRow("SELECT COUNT(1) FROM Orders").Scan(&count)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, count)
}

func (suite *OrderRepositorySuite) TestOrderRepository_GetTotal() {
	var total int
	var err error

	// test with empty table
	total, err = suite.OrderRepo.GetTotal()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, total)

	// test with non-empty table
	order1 := &entity.Order{
		ID:         "1",
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}

	order2 := &entity.Order{
		ID:         "2",
		Price:      200,
		Tax:        20,
		FinalPrice: 220,
	}

	err = suite.OrderRepo.Save(order1)
	assert.NoError(suite.T(), err)

	err = suite.OrderRepo.Save(order2)
	assert.NoError(suite.T(), err)

	total, err = suite.OrderRepo.GetTotal()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, total)
}

func TestOrderRepositorySuite(t *testing.T) {
	suite.Run(t, new(OrderRepositorySuite))
}