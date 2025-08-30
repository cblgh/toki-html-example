// THIS IS GENERATED, PLS NO EDITS
package translations
import (
	"eyeneighteenn/tokibundle"
	"time"
)
func Translations(t tokibundle.Reader) {
	time.Now()
	// LOOP
	t.String("Flour site")
	t.String("My flour lists")
	t.String("On this page I will collect a few of my favourite flour varieties")
	t.String("Gluten-free flours")
	t.String("Gluten-free flours are important for allergy reasons and being able to make things for friends that won't kill them. They are also important for bread-baking.")
	t.String("When baking gluten-free goods, it's typically best to mix a few different kinds of flours")
	t.String("Rice flour")
	t.String("Chickpea flour")
	t.String("Almond flour")
	t.String("Wheat flours")
	t.String("When it comes to wheat flours, there is a great variety and with varying levels of nutrition. The more of the grain left in the flour, the better. Stone-milled flours are not just a fad, but a sign that says that this flour likely has more fiber and nutrition")
	t.String("Spelt flour")
	t.String("Graham flour")
	t.String("Spring wheat flour (aka bread flour)")
	t.String("{date-medium} {name} proposed: {text}", time.Now(), tokibundle.String{Value: "Neutralo", Gender: tokibundle.GenderNeutral}, "GenericAction")
}