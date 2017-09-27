package main

import (
	e "alchill/alcohol_data"
	"io/ioutil"
	"os"
)

var lvl1 e.Spirits

func main() {

	f1, _ := ioutil.ReadFile("./alcohol_data/1.json")
	lvl1.UnmarshalJSON(f1)
	f2, _ := ioutil.ReadFile("./alcohol_data/2.json")
	f3, _ := ioutil.ReadFile("./alcohol_data/3.json")
	f4, _ := ioutil.ReadFile("./alcohol_data/4.json")

	lvl1 = createJSONWithBaseCategories(f1)
	lvl1 = createJSONWithBaseCategories(f2)
	lvl1 = createJSONWithBaseCategories(f3)
	lvl1 = createJSONWithBaseCategories(f4)

	f, _ := os.Create("output.json")
	defer f.Close()
	temp, _ := lvl1.MarshalJSON()
	f.Write(temp)

}

func createJSONWithBaseCategories(file []byte) e.Spirits{
	next_lvl := e.SpiritsWOCategory{}
	next_lvl.UnmarshalJSON(file)
	next_lvl_result := []e.Spirit{}
	for _, obj := range next_lvl{
		result := e.Spirit{}
		result.ID = obj.ID
		result.Name = obj.Name
		result.IsCategory = obj.IsCategory
		result.Bars = obj.Bars
		result.Caloricity = obj.Caloricity
		result.Value = obj.Value
		result.Volume = obj.Volume

		if len(next_lvl[0].Categories) == 1 {
			for _, o := range lvl1{
				if obj.Categories[0] == o.ID {
					result.Categories = append(result.Categories, o)
				}
			}
			next_lvl_result = append(next_lvl_result, result)
		}
		if len(next_lvl[0].Categories) == 2 {
			for _, o := range lvl1{
				if obj.Categories[0] == o.ID {
					result.Categories = append(result.Categories, o)
				}
				if obj.Categories[1] == o.ID {
					result.Categories = append(result.Categories, o)
				}
			}
			next_lvl_result = append(next_lvl_result, result)
		}
		if len(next_lvl[0].Categories) == 3 {
			for _, o := range lvl1{
				if obj.Categories[0] == o.ID {
					result.Categories = append(result.Categories, o)
				}
				if obj.Categories[1] == o.ID {
					result.Categories = append(result.Categories, o)
				}
				if obj.Categories[2] == o.ID {
					result.Categories = append(result.Categories, o)
				}
			}
			next_lvl_result = append(next_lvl_result, result)
		}
	}
	return append(lvl1, next_lvl_result...)
}