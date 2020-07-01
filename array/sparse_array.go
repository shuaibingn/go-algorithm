package array

type sparseNode struct {
	row, col int
	val      interface{}
}

type sparseArray struct {
	data                 []sparseNode
	lengthRow, lengthCol int
}

func NewSparseArray(originArray [][]interface{}) *sparseArray {
	var sparseArr []sparseNode
	for i1, v1 := range originArray {
		for i2, v2 := range v1 {
			sparseNodeTmp := sparseNode{
				row: i1,
				col: i2,
				val: v2,
			}
			sparseArr = append(sparseArr, sparseNodeTmp)
		}
	}
	return &sparseArray{
		data:      sparseArr,
		lengthRow: len(originArray),
		lengthCol: len(originArray[0]),
	}
}

func (s *sparseArray)TransToArray() [][]interface{}{

	originArr := make([][]interface{},  s.lengthRow)
	for k, _ := range originArr {
		resultArr := make([]interface{}, s.lengthCol)
		originArr[k] = resultArr
	}

	for _, v := range s.data {
		originArr[v.row][v.col] = v.val
	}
	return originArr
}
