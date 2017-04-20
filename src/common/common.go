package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"strings"
	"time"
)

type R struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type RE struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func ReturnFormat(w http.ResponseWriter, code int64, data interface{}, msg string) {
	res := R{Code: code, Data: data, Msg: msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}

func ReturnEFormat(w http.ResponseWriter, code int64, msg string) {
	res := RE{Code: code, Msg: msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}

func UnmarshalSheet(sheet *xlsx.Sheet, v interface{}) (err error) {
	fieldsMap := make(map[int][2]string)
	data := []map[string]interface{}{}

	for index, row := range sheet.Rows {
		if index == 0 {
			for i, cell := range row.Cells {
				if strings.Contains(cell.Value, ":") {
					a := strings.Split(cell.Value, ":")

					fieldsMap[i] = [2]string{a[0], a[1]}
				} else {
					fieldsMap[i] = [2]string{cell.Value, ""}
				}
			}
			continue
		}

		rowMap := map[string]interface{}{}
		arrMap := map[string][]map[string]string{}
		var cv string

		// 列数为0跳过
		if len(row.Cells) == 0 {
			continue
		}

		for i, cell := range row.Cells {
			if len(fieldsMap[i]) == 2 {
				cv, err = cell.SafeFormattedValue()
				if err != nil {
					err = fmt.Errorf("%s解析失败: %s", fieldsMap[i][0], err.Error())
					return
				}

				// EXECL 时间特殊处理
				// EXECL 解析后的时间对应的 CellTypeNumeric 类型!!!
				if cell.Type() == xlsx.CellTypeNumeric {
					// 修证EXECL数据 2016\-12\-01 为 2016-12-01
					cv = strings.Replace(cv, "\\", "", -1)

					// 修证EXECL数据 2016/12/01 为 2016-12-01
					cv = strings.Replace(cv, "/", "-", -1)
				}

				// 包含2个"-"、2个":"则看下是否是时间格式
				if strings.Count(cv, "-") == 2 && strings.Count(cv, ":") == 2 {
					_, err2 := ParseDatetime(cv, "2006-01-02 15:04:05")
					if err2 == nil {
						cv += " +0800"
					}
				}

				if k, v := fieldsMap[i][0], fieldsMap[i][1]; v != "" {
					if len(arrMap[k]) == 0 {
						arrMap[k] = []map[string]string{}
					}

					tmp := map[string]string{}
					tmp[v] = cv

					arrMap[k] = append(arrMap[k], tmp)
					rowMap[k] = arrMap[k]
				} else {
					rowMap[k] = cv
				}
			}
		}

		if len(rowMap) > 0 {
			data = append(data, rowMap)
		}
	}

	d, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(d, v)
	return
}

func ParseDatetime(str, format string) (datetime time.Time, err error) {
	if str != "" {
		datetime, err = time.ParseInLocation(format, str, time.Now().Location())

		if err != nil || datetime.IsZero() {
			err = errors.New(fmt.Sprintf("datetime format error. except: %s, but got %s", format, str))
		}
	} else {
		err = errors.New("nil")
	}
	return
}
