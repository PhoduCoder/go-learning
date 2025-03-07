point2 := struct {
    x int
    y int
  }{}
  point2.x = 10
  point2.y = 5

Here point2 is an anonymous struct 
Note that there is no name for this struct like a struct generally has 

point1 := struct {
    x int
    y int
  }{
    10,
    10,
  }

  Another way of declaring anonymous struct 
========
type point struct {
	x int
	y int
  }

point3 := point{10, 10}

Normal struct to compare this with anonymous one 