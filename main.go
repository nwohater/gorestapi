package main

// tutorial: https://www.youtube.com/watch?v=d_L64KT3SFM
// rest api using gin framework

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "First Item", Completed: false},
	{ID: "2", Item: "Second Item", Completed: false},
	{ID: "2", Item: "Third Item", Completed: false},
}
