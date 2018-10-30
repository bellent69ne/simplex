package util

import (
    "fmt"
)

func calcCanProceed(table [][]float64, inverted bool) bool {
    if !inverted {
        for _, val := range table[len(table) - 1] {
            if val < 0 {
                return true
            }
        }
    } else {
        for _, row := range table {
            if row[len(row) - 1] < 0 {
                return true
            }
        }
    }

    return false
}

func printTable(table [][]float64) {
    fmt.Println()
    for _, row := range table {
        for _, val := range row {
            fmt.Printf("%.2f\t\t", val)
        }
        fmt.Println()
    }
    fmt.Println()
}

func Calculate(table [][]float64, inverted bool) {
    printTable(table)
    for calcCanProceed(table, inverted) {
        row, column := pivotElem(table)
        makePivElemOne(table, row, column)
        printTable(table)
        makeZeroUnderPivElem(table, row, column)
        printTable(table)

    }
}

func pivotElem(table [][]float64) (row, column int) {
    min := 0.0
    for i, val := range table[len(table) - 1] {
        if val < min {
            min = val
            column = i
        }
    }
    //fmt.Println("Column ", column)

    min = 10000000000.0

    for i := 0; i < len(table) - 1; i++ {
        someVar := table[i][len(table[i]) - 1] / table[i][column]
//        fmt.Println("someVar ", someVar)
        if someVar < min {
            min = someVar
            row = i
        }
    }

   // fmt.Println("Min ", table[row][column])
  //  fmt.Println("indexes: ", row, column)
    return
}

func makePivElemOne(table [][]float64, row, column int) {
    divisor := table[row][column]
    for i := 0; i < len(table[row]); i++ {
        table[row][i] *= (1 / divisor)
    }
}

func makeZeroUnderPivElem(table [][]float64, row, column int) {
    if row > 0 {
        for i := 0; i < row; i++ {
            val := table[i][column]
            for j := 0; j < len(table[i]); j++ {
                table[i][j] -= val * table[row][j]
            }
        }
    }
    for i := 1; i < len(table); i++ {
        if row + i >= len(table) {
            break
        }
        val := table[row + i][column]
        for j := 0; j < len(table[row + i]); j++ {
            table[row + i][j] -= val * table[row][j]
        }
    }
}
