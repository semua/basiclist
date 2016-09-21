package basiclist

import (
	"fmt"
	"testing"
	"time"
)

func Compare(val0 interface{}, val1 interface{}) int {
	pubt0 := val0.(Article).pubtime - val0.(Article).pubtime%(3600*12)
	pubt1 := val1.(Article).pubtime - val1.(Article).pubtime%(3600*12)
	if pubt0-pubt1 == 0 {
		if val0.(Article).score-val1.(Article).score > 0 {
			return 1
		}
		if val0.(Article).score-val1.(Article).score < 0 {
			return -1
		}
		return 0
	}
	if pubt0-pubt1 > 0 {
		return 1
	}
	return -1
}
func CompareDesc(val0 interface{}, val1 interface{}) int {
	pubt0 := val0.(Article).pubtime - val0.(Article).pubtime%(3600*12)
	pubt1 := val1.(Article).pubtime - val1.(Article).pubtime%(3600*12)
	if pubt0-pubt1 == 0 {
		if val0.(Article).score-val1.(Article).score > 0 {
			return -1
		}
		if val0.(Article).score-val1.(Article).score < 0 {
			return 1
		}
		return 0
	}
	if pubt0-pubt1 > 0 {
		return -1
	}
	return 1
}

type Article struct {
	title   string
	pubtime int64
	score   float64
}

func TestBasicList(t *testing.T) {
	bList := NewBasicList()

	art1 := Article{title: "title1", pubtime: time.Now().Unix() - 1000, score: 1.704}
	art2 := Article{title: "title2", pubtime: time.Now().Unix() - 1100, score: 1.204}
	art3 := Article{title: "title3", pubtime: time.Now().Unix() - 1500, score: 1.804}
	art4 := Article{title: "title4", pubtime: time.Now().Unix() - 1100, score: 1.504}
	art5 := Article{title: "title5", pubtime: time.Now().Unix() - 110000, score: 2.504}

	bList.Insert(Compare, "title1", art1)
	bList.Print()
	bList.Insert(Compare, "title2", art2)
	bList.Print()
	bList.Insert(Compare, "title3", art3)
	bList.Print()
	bList.Insert(Compare, "title4", art4)
	bList.Print()
	bList.Insert(Compare, "title5", art5)
	fmt.Println(bList.size)
	bList.Print()
	fmt.Println(bList.GetByKey("title2"))
	fmt.Println(bList.GetByKey("titlen"))

	fmt.Println(bList.Equal(Compare, art1))
	art1_1 := Article{title: "title11", pubtime: time.Now().Unix() - 1000, score: 1.700}
	fmt.Println(bList.Equal(Compare, art1_1))
	art1_2 := Article{title: "title12", pubtime: time.Now().Unix() - 1000, score: 1.5}
	fmt.Println(bList.Equal(Compare, art1_2))
	art1_3 := Article{title: "title13", pubtime: time.Now().Unix() - 1000, score: 0.5}
	fmt.Println(bList.Equal(Compare, art1_3))
	art1_4 := Article{title: "title14", pubtime: time.Now().Unix() - 1000, score: 5.5}
	fmt.Println(bList.Equal(Compare, art1_4))
	art1_5 := Article{title: "title15", pubtime: time.Now().Unix() - 100000, score: 0.5}

	cList := NewBasicList()
	fmt.Println(cList.Equal(Compare, art1_1))
	fmt.Println(bList.Equal(Compare, art2))
	fmt.Println(bList.Equal(Compare, art3))
	fmt.Println(bList.Equal(Compare, art4))
	fmt.Println(bList.Equal(Compare, art5))
	fmt.Println("LessThanOrEqual")
	fmt.Println(bList.LessThanOrEqual(Compare, art1))
	fmt.Println(bList.LessThanOrEqual(Compare, art2))
	fmt.Println(bList.LessThanOrEqual(Compare, art3))
	fmt.Println(bList.LessThanOrEqual(Compare, art4))
	fmt.Println(bList.LessThanOrEqual(Compare, art5))
	fmt.Println("LessThanOrEqual1")
	fmt.Println(bList.LessThanOrEqual(Compare, art1_1))
	fmt.Println(bList.LessThanOrEqual(Compare, art1_2))
	fmt.Println(bList.LessThanOrEqual(Compare, art1_3))
	fmt.Println(bList.LessThanOrEqual(Compare, art1_4))
	fmt.Println(bList.LessThanOrEqual(Compare, art1_5))
	fmt.Println(cList.LessThanOrEqual(Compare, art1_5))
	fmt.Println("GreaterThanOrEqual")
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art2))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art3))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art4))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art5))
	fmt.Println("GreaterThanOrEqual1")
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1_1))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1_2))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1_3))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1_4))
	fmt.Println(bList.GreaterThanOrEqual(Compare, art1_5))
	fmt.Println(cList.GreaterThanOrEqual(Compare, art1_5))
	fmt.Println("List")
	fmt.Println(bList.List(0, 2))
	fmt.Println(bList.List(1, 3))
	fmt.Println(bList.List(1, 0))
	fmt.Println(bList.List(3, 5))
	fmt.Println(bList.Remove("titlen"))
	fmt.Println(bList.Remove("title4"))
	bList.Print()
	fmt.Println(bList.Remove("title5"))
	bList.Print()
	fmt.Println(bList.Remove("title2"))
	bList.Print()
	fmt.Println(bList.Remove("title3"))
	bList.Print()
	fmt.Println(bList.Remove("title1"))
	bList.Print()
}
func TestBasicListDesc(t *testing.T) {
	bList := NewBasicList()

	art1 := Article{title: "title1", pubtime: time.Now().Unix() - 1000, score: 1.704}
	art2 := Article{title: "title2", pubtime: time.Now().Unix() - 1100, score: 1.204}
	art3 := Article{title: "title3", pubtime: time.Now().Unix() - 1500, score: 1.804}
	art4 := Article{title: "title4", pubtime: time.Now().Unix() - 1100, score: 1.504}
	art5 := Article{title: "title5", pubtime: time.Now().Unix() - 110000, score: 2.504}

	bList.Insert(CompareDesc, "title1", art1)
	bList.Print()
	bList.Insert(CompareDesc, "title2", art2)
	bList.Print()
	bList.Insert(CompareDesc, "title3", art3)
	bList.Print()
	bList.Insert(CompareDesc, "title4", art4)
	bList.Print()
	bList.Insert(CompareDesc, "title5", art5)
	fmt.Println(bList.size)
	bList.Print()
	fmt.Println(bList.GetByKey("title2"))
	fmt.Println(bList.GetByKey("titlen"))

	fmt.Println(bList.Equal(CompareDesc, art1))
	art1_1 := Article{title: "title11", pubtime: time.Now().Unix() - 1000, score: 1.700}
	fmt.Println(bList.Equal(CompareDesc, art1_1))
	art1_2 := Article{title: "title12", pubtime: time.Now().Unix() - 1000, score: 1.5}
	fmt.Println(bList.Equal(CompareDesc, art1_2))
	art1_3 := Article{title: "title13", pubtime: time.Now().Unix() - 1000, score: 0.5}
	fmt.Println(bList.Equal(CompareDesc, art1_3))
	art1_4 := Article{title: "title14", pubtime: time.Now().Unix() - 1000, score: 5.5}
	fmt.Println(bList.Equal(CompareDesc, art1_4))
	art1_5 := Article{title: "title15", pubtime: time.Now().Unix() - 100000, score: 0.5}

	cList := NewBasicList()
	fmt.Println(cList.Equal(CompareDesc, art1_1))
	fmt.Println(bList.Equal(CompareDesc, art2))
	fmt.Println(bList.Equal(CompareDesc, art3))
	fmt.Println(bList.Equal(CompareDesc, art4))
	fmt.Println(bList.Equal(CompareDesc, art5))
	fmt.Println("LessThanOrEqual")
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art2))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art3))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art4))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art5))
	fmt.Println("LessThanOrEqual1")
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1_1))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1_2))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1_3))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1_4))
	fmt.Println(bList.LessThanOrEqual(CompareDesc, art1_5))
	fmt.Println(cList.LessThanOrEqual(CompareDesc, art1_5))
	fmt.Println("GreaterThanOrEqual")
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art2))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art3))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art4))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art5))
	fmt.Println("GreaterThanOrEqual1")
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1_1))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1_2))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1_3))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1_4))
	fmt.Println(bList.GreaterThanOrEqual(CompareDesc, art1_5))
	fmt.Println(cList.GreaterThanOrEqual(CompareDesc, art1_5))
	fmt.Println("List")
	fmt.Println(bList.List(0, 2))
	fmt.Println(bList.List(1, 3))
	fmt.Println(bList.List(1, 0))
	fmt.Println(bList.List(3, 5))
	fmt.Println(bList.Remove("titlen"))
	fmt.Println(bList.Remove("title4"))
	bList.Print()
	fmt.Println(bList.Remove("title5"))
	bList.Print()
	fmt.Println(bList.Remove("title2"))
	bList.Print()
	fmt.Println(bList.Remove("title3"))
	bList.Print()
	fmt.Println(bList.Remove("title1"))
	bList.Print()
}
