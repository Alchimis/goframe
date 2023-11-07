package dbs

import (
	"errors"
	"fmt"
	internal "inter"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
	c "github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	d "github.com/ostafen/clover/v2/document"
	q "github.com/ostafen/clover/v2/query"
)

type CloverDB struct {
	Destination string
	Db          *c.DB
	Colections  []string
}

func Init(colections []string) *CloverDB {
	return &CloverDB{
		Destination: "c:\\dbs",
		Colections:  colections,
	}
}

// возможно часть интерфейса
func (db *CloverDB) GetRecordWhere(from string, conditions map[string]interface{}) (*d.Document, error) {
	searchQuery := q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		for key := range conditions {
			if doc.Get(key) == nil {
				return false
			}
		}
		return true
	})
	return db.Db.FindFirst(searchQuery)
}

// возможно часть интерфейса
func (db *CloverDB) InsertRecord(to string, record map[string]interface{}) (string, error) {
	doc := d.NewDocument()
	for val, key := range record {
		doc.Set(val, key)
	}
	return db.Db.InsertOne(to, doc)
}

func (db *CloverDB) Run() (*CloverDB /* или *Database*/, error) {
	var err error
	db.Db, err = c.Open(db.Destination)
	if err != nil {
		return nil, err
	}
	for _, val := range db.Colections {
		if err := db.Db.CreateCollection(val); err != nil {
			if errors.Is(err, c.ErrCollectionExist) {
				log.Println("Colection "+val+" already exists: ", err)

			} else {
				//log.Println("Error with colection creation", err)
				return nil, err
			}
		}
	}
	return db, nil
}

func DocsToUsers(docs []*d.Document) ([]*internal.User, error) {
	log.Println("Docks: ", docs)
	log.Println("Dock len: ", len(docs))
	output := make([]*internal.User, len(docs))

	var err error
	for key, val := range docs {
		usr := &internal.User{}
		err = val.Unmarshal(usr)
		if err != nil {
			hhh, _ := uuid.ParseBytes(val.Get("id").([]byte))
			log.Println("terpi: ", hhh)
			log.Println("Error with unmarshal: ", err)
			log.Println("val: ", val)
			log.Println("ya uze: ", val.Get("id"))
			log.Println("ya yobnulsa: ", string(val.Get("id").([]byte)))
			//	return nil, err
		}
		output[key] = usr
	}
	return output, nil
}

func PrintDocs(docs []*d.Document) {
	for _, v := range docs {
		log.Print(*v, " ")
	}
}

// TODO: если ты и сел ебланить то делай замеры блять
func (db *CloverDB) GetAllUsers(from string) ([]*internal.User, error) {
	//var tim time.Time
	users, err := db.Db.FindAll(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		return doc.Has("nickname") && doc.Has("password") // && doc.Has("live_connections")
	}))
	PrintDocs(users)
	if err != nil {
		return nil, err
	}
	/*output := make([]internal.User, len(users))
	var usr *internal.User
	for key, val := range users {
		err = val.Unmarshal(usr)
		if err != nil {
			return nil, err
		}
		output[key] = *usr
	}*/
	output, err := DocsToUsers(users)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func Run() {
	db, err := c.Open("c:\\dbs")
	if err != nil {
		log.Println("Error with db creation: ", err)
	}
	fmt.Println(db)
	time.Sleep(time.Second * 30)
}

func UserToRecord(user *internal.User) map[string]interface{} {
	record := make(map[string]interface{}, 0)
	record["nickname"] = user.Nickname
	record["password"] = user.Password
	record["id"] = user.Id
	record["join_at"] = user.JoinAt
	record["live_conections"] = user.LiveConnections
	return record
}

func ChannelToRecord(chanel *internal.Channel) map[string]interface{} {
	record := make(map[string]interface{}, 0)
	record["id"] = chanel.Id
	record["name"] = chanel.Name
	record["created_at"] = chanel.CreatedAt
	return record
}

type ErrNoSuchField struct {
	field string
}

func (err ErrNoSuchField) Error() string {
	return "No such field: " + err.field
}

type ErrFieldNotThatType struct {
	expectedType, field string
}

func (err ErrFieldNotThatType) Error() string {
	return "Field " + err.field + " not type of " + err.expectedType
}

type ErrNotTypeOf struct {
	expextedtType, variable string
}

func (err ErrNotTypeOf) Error() string {
	return "Variable " + err.variable + " is no type of " + err.expextedtType
}

func InterfaceToUUID(mb interface{}) (uuid.UUID, error) {
	bt, ok := mb.([]byte)
	if !ok {
		return uuid.Nil, ErrFieldNotThatType{"byte", "id"}
	}
	id, err := uuid.ParseBytes(bt)
	if err != nil {
		return uuid.Nil, err
	}
	return id, err
}

func IntarfeceTo[T any](mb interface{}) (*T, error) {
	target, ok := mb.(T)
	if !ok {
		return nil, ErrNotTypeOf{"target", "string"}
	}
	return &target, nil
}

func InterfaceToString(mb interface{}) (string, error) {
	target, ok := mb.(string)
	if !ok {
		return "", ErrNotTypeOf{"target", "string"}
	}
	return target, nil
}

func ChanelToDoc(chn *internal.Channel) (*document.Document, error) {
	doc := document.NewDocumentOf(chn)
	err := doc.Unmarshal(chn)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type Focus struct {
	val interface{}
}

func (db *CloverDB) GetAllFocusUser(from string) (interface{}, error) {
	//var tim time.Time
	return db.Db.FindAll(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		return (doc.Has("nickname") || doc.Has("Nickname")) && (doc.Has("password") || doc.Has("Password")) // && doc.Has("live_connections")
	}))

}

// я нихуево так наебланил. Не думаю вообще. Для нормальных бд пишут блядь скрипты. Тут такая же ситуация. Только бля проблема в том что я еблан.
// ЕСЛИ ЧЁ УСЛОВИЯ МОЖНО ПЕРЕДАВАТЬ ЧЕРЕЗ ИНТЕРФЕЙС И КАСТИТЬ ИХ ВК ФУНЦИЯМ СКРИПТАМ И ПРОЧЕЙ ХУЕТЕ

func (db *CloverDB) GetAllFocusUserWhere(from string, conditions map[string]interface{}) (interface{}, error) {
	//var tim time.Time
	return db.Db.FindAll(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		if (doc.Has("nickname") || doc.Has("Nickname")) && (doc.Has("password") || doc.Has("Password")) {
			for key := range conditions {
				if !reflect.DeepEqual(doc.Get(key), conditions[key]) {
					return false
				}
				return true
			}
		}
		return false
	}))
}

func (db *CloverDB) GetAllInterfaceUserWhere(from string, inter interface{}) (interface{}, error) {
	//var tim time.Time
	conditions, ok := inter.(func(*d.Document) bool)
	if !ok {
		return nil, errors.New("Eblan, ti peredal ne functiu dolbayob")
	}
	return db.Db.FindAll(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		if !(doc.Has("nickname") || doc.Has("Nickname")) && (doc.Has("password") || doc.Has("Password")) {
			return false
		}
		return conditions(doc)
	}))
}

// TODO: разобраться с сериализаций времени
func (db *CloverDB) InsertFocusRecord(to string, focus interface{}) (string, error) {
	//doc := focus.(*d.Document)
	//log.Println("docdoc: ", doc)
	return db.Db.InsertOne(to, focus.(*d.Document))
}

func (db *CloverDB) DeleteAllWhere(from string, conditions map[string]interface{}) (interface{}, error) {
	err := db.Db.Delete(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		for key := range conditions {
			if !(reflect.DeepEqual(doc.Get(key), conditions[key])) {
				return false
			}
		}
		return true
	}))
	return nil, err
}

// я могу нанизывать условия. если я расширю интерфейс до двух функций я смогу и условия дополнять. к сожалению не особо гибко получаеться
// можно сделать настройки и передавать всякие скрипты к какую то отдельную структуру.
//
//	это какой то рофл. мне надо имет какой то интерфейс скрипта. причём для всех скриптов чтобы это всё дело было достаточно гибким.
//
// проблема в том что я взял простую бд специально чтобы всё сдклать побыстрее, а в итоге полез в залупу.
// в реляционных бд есть так сказать таблицы и поля. в таких как кловер есть только файлы, так что нужно отделять их както
// можно в качестве инициализации для поиска запросов использовать структуру, потом уже к ней добавлять всякие условия поиска
// мне не особо понятно как в не реаляционных бд работает слияние таблиц. если между ними разница будет огромной тогда я блять не ебу
// я думаю разница между ними будет не такая огромная, но реализовывать общий шаблон для запросов я СЕЙЧАС НЕ БУДУ
// я вывел для себя пока алгоритм, который не будет гибким.
// его можно использовать как уменьшение кода для писания кода под капотам таких запросов как получение юзера по id
// и вообще я как еблан проебал дохуя времени, ниписав нихуя
func (db *CloverDB) DeleteAllInterfaceWhere(from string, focus interface{}) (interface{}, error) {
	conditions, ok := focus.(func(*d.Document) bool)
	if !ok {
		return nil, errors.New("Eblan, ti peredal ne functiu dolbayob")
	}
	err := db.Db.Delete(q.NewQuery(from).MatchFunc(func(doc *d.Document) bool {
		return conditions(doc)
	}))
	return nil, err
}

func (db *CloverDB) GetChanelByIdRaw(id string) (*d.Document, error) {
	return db.Db.FindFirst(q.NewQuery("channels").MatchFunc(
		func(doc *d.Document) bool {
			if !(doc.Has("id") && doc.Has("cretaed_at") && doc.Has("users")) {
				return false
			}
			idd := doc.Get("id")
			iddd, ok := idd.(string)
			if !ok {
				return false
			}
			return iddd == id
		},
	))
}
func (db *CloverDB) GetChanelById(id string) (*internal.Channel, error) {
	doc, err := db.GetChanelByIdRaw(id)
	if err != nil {
		return nil, err
	}
	b := &struct {
		Id        string    `clover:"id,string"`
		Name      string    `clover:"name"`
		CreatedAt time.Time `clover:"cretaed_at,string,omitempty"`
		Users     []string  `clover:"users"`
	}{}
	err = doc.Unmarshal(b)
	if err != nil {
		return nil, err
	}
	return &internal.Channel{
		Id:        b.Id,
		Name:      b.Name,
		Users:     b.Users,
		CreatedAt: b.CreatedAt,
	}, nil
}
