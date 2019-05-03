package context

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
	"github.com/stampjohnny/mttv/utils"
)

type Context struct {
	cryptoAmount float64
	moneyAmount  float64
	trxn         uuid.UUID
	price        float64
}

var db *badger.DB
var currentContext *Context

func (c *Context) GetTransactionID() uuid.UUID {
	return c.trxn
}

func (c *Context) SetTransactionID(trn uuid.UUID) {
	c.trxn = trn
}

func (c *Context) GetCryptoAmount() float64 {
	return c.cryptoAmount
}

func (c *Context) SetCryptoAmount(value float64) {
	c.cryptoAmount = value
}

// func (c *Context) GetMoneyAmount() string {
// 	return govalidator.ToString(c.moneyAmount)
// }

// func (c *Context) SetMoneyAmount(value interface{}) error {
// 	val, err := getHybridFloat(value)
// 	if err != nil {
// 		return fmt.Errorf("can't set crypto amount: %s", err)
// 	}
// 	c.moneyAmount = float64(val)
// 	return nil
// }

// func (c *Context) GetCryptoPrice() string {
// 	return govalidator.ToString(c.price)
// }

// func (c *Context) SetCryptoPrice(price interface{}) error {
// 	value := govalidator.ToString(price)
// 	var err error
// 	c.price, err = govalidator.ToFloat(value)
// 	if err != nil {
// 		return fmt.Errorf("can't convert price to float: %v %s", price, err)
// 	}
// 	return nil
// }

func Init() error {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	badgetDB, err := badger.Open(opts)
	utils.Crash(err != nil, fmt.Sprintf("can't initialize context: %s", err))
	db = badgetDB
	return nil
}

func validate() error {
	if currentContext.GetTransactionID().String() == "" {
		return errors.New("validation error: transaction ID is empty")
	}
	if currentContext.GetCryptoAmount() <= 0 {
		return errors.New("validation error: crypt amount is empty")
	}
	// if currentContext.GetMoneyAmount() == "" {
	// 	return errors.New("validation error: money amount is empty")
	// }
	// if currentContext.GetCryptoPrice() == "" {
	// 	return errors.New("validation error: crypto price is empty")
	// }
	return nil
}

func New() *Context {
	return &Context{}
}

func Get() *Context {
	return currentContext
}

func reset() {
	currentContext = New()
}

func Find(trn string) error {
	reset()
	txn := db.NewTransaction(true)
	defer txn.Discard()

	key := createDbKey(trn, "ca")
	fmt.Printf("find key %+v\n", key) // output for debug
	item, err := txn.Get(key)
	if err != nil {
		return fmt.Errorf("can't find crypto amount: %s", err)
	}
	valueBytes, err := item.ValueCopy(nil)
	if err != nil {
		return fmt.Errorf("can't find crypto amount: %s", err)
	}
	value := string(valueBytes)
	fmt.Printf("find value %+v\n", value) // output for debug

	cryptoAmount, err := govalidator.ToFloat(value)
	if err != nil {
		return fmt.Errorf("can't convert crypto amount to float: %s", err)
	}
	currentContext.SetCryptoAmount(cryptoAmount)

	uid, err := uuid.Parse(trn)
	if err != nil {
		return fmt.Errorf("can't parse uuid: %s", err)
	}
	currentContext.SetTransactionID(uid)

	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

func Set(context interface{}) {
	currentContext = context.(*Context)
}

func Save() error {
	utils.Crash(db == nil, "database is not initialized")

	if err := validate(); err != nil {
		return fmt.Errorf("can't save context: %s", err)
	}
	txn := db.NewTransaction(true)
	defer txn.Discard()

	trn := currentContext.GetTransactionID()

	key := createDbKey(trn.String(), "ca")
	fmt.Printf("save key %+v\n", key)                                // output for debug
	fmt.Printf("save value %+v\n", currentContext.GetCryptoAmount()) // output for debug
	err := txn.Set(key, toBytes(currentContext.GetCryptoAmount()))
	if err != nil {
		return fmt.Errorf("can't save transaction id: %s", err)
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
func createDbKey(trn string, sufix string) []byte {
	return []byte(fmt.Sprintf("%s.%s", trn, sufix))
}

func toBytes(value interface{}) []byte {
	return []byte(govalidator.ToString(value))
}
