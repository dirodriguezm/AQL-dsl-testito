package core

import "context"

type Object struct {
	Oid     string  `json:"oid"`
	Meanra  float32 `json:"meanra"`
	Meandec float32 `json:"meandec"`
}

type ObjectUseCase struct{}

func makeObjectQuery(queryArgs *QueryArguments) string {
	query := ""
	if queryArgs.Oid != "" {
		query = "SELECT oid, meanra, meandec from object where oid=$1"
	}
	return query
}

func (uc *ObjectUseCase) GetObject(queryArgs *QueryArguments) []Object {
	sqlQuery := makeObjectQuery(queryArgs)
	conn := getDbConn()
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), sqlQuery, queryArgs.Oid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var objects []Object
	for rows.Next() {
		var object Object
		err := rows.Scan(&object.Oid, &object.Meanra, &object.Meandec)
		if err != nil {
			panic(err)
		}
		objects = append(objects, object)
	}
	return objects
}
