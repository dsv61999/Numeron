module example.com/numeron

go 1.15

replace numeron/pkg => ./pkg

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2
	numeron/pkg v0.0.0-00010101000000-000000000000
)
