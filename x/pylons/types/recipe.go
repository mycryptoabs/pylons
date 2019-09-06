package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

// Recipe is a game state machine step abstracted out as a cooking terminology
type Recipe struct {
	CookbookName  string // the cookbook guid
	RecipeName    string
	ID            string // the recipe guid
	CoinInputs    CoinInputList
	CoinOutputs   CoinOutputList
	ItemInputs    ItemInputList
	ItemOutputs   ItemOutputList
	Description   string
	ExecutionTime int64
	Sender        sdk.AccAddress
	Disabled      bool
}

// RecipeList is a list of cookbook
type RecipeList struct {
	Recipies []Recipe
}

func (cbl RecipeList) String() string {
	output := "RecipeList{"
	for _, cb := range cbl.Recipies {
		output += cb.String()
		output += ",\n"
	}
	output += "}"
	return output
}

// NewRecipe creates a new recipe
func NewRecipe(recipeName, cookbookName, description string,
	coinInputs CoinInputList, coinOutputs CoinOutputList, itemInputs ItemInputList, itemOutputs ItemOutputList,
	execTime int64, sender sdk.AccAddress) Recipe {
	rcp := Recipe{
		RecipeName:    recipeName,
		CookbookName:  cookbookName,
		CoinInputs:    coinInputs,
		CoinOutputs:   coinOutputs,
		ItemInputs:    itemInputs,
		ItemOutputs:   itemOutputs,
		ExecutionTime: execTime,
		Description:   description,
		Sender:        sender,
	}

	rcp.ID = rcp.KeyGen()
	return rcp
}

func (rcp *Recipe) String() string {
	return fmt.Sprintf(`Recipe{
		RecipeName: %s,
		CookbookName: %s,
		ID: %s,
		CoinInputs: %s,
		CoinOutputs: %s,
		ItemInputs: %s,
		ItemOutputs: %s,
		ExecutionTIme: %d,
	}`, rcp.RecipeName, rcp.CookbookName, rcp.ID, rcp.CoinInputs.String(),
		rcp.CoinOutputs.String(), rcp.ItemInputs.String(), rcp.ItemOutputs.String(), rcp.ExecutionTime)
}

// KeyGen generates key for the store
func (rc Recipe) KeyGen() string {
	id := uuid.New()
	return rc.Sender.String() + id.String()
}
