package main

import (
    "github.com/ballent69ne/simplex/util"
)

func main() {
    //simplexTable := [][]float64{{2,   1,  1, 1, 0, 0, 0, 180},
      //                          {1,   3,  2, 0, 1, 0, 0, 300},
        //                        {2,   1,  2, 0, 0, 1, 0, 240},
          //                      {-6, -5, -4, 0, 0, 0, 1, 0  }}
    simplexTable := [][]float64{{0.2,  0.75,  0, 1, 0, 0, 2},
                                {0.1,  0.1, 0.1, 0, 1, 0, 3},
                                {-120, -70, -10, 0, 0, 1, 0}}
    util.Calculate(simplexTable, false)
}
