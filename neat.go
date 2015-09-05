package main

type Genome struct {
  conns []Conn
}

type Conn struct {
  inNode    Node 
  outNode   Node
  weight    int
  expressed bool
  innov_num int
}

type Node struct {
  is_input  Bool
  is_hidden Bool
  is_output Bool
}

func New( c Conn ) Conn {
  
}

//
// Single new connection gene with random weight added connecting
// two previously unconnected nodes
//
func add_conn( node1 Node, node2 Node ) Conn {
  weight         := 1
  innovation_num := 0
  conn           := Conn{ node1, node2, weight, true, innovation_num } 
  return( conn )
}

//
// Existing connection is split and new node placed where the connection was
// Old conn disabled, two new conns created between old end points to new node
//
func add_node( conn Conn ){
  conn.expressed = false

  innovation_num := 0
  node  := Node{ is_hidden : true }
  conn1 := Conn{ conn.inNode, node, 1, true, innovation_num }
  conn2 := Conn{ node, conn.outNode, conn.weight, true, innovation_num }
}

func main() {
  innovation_num := 0
}
