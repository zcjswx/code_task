package app

type Tabler interface {
	TableName() string
}

type GraphEntity struct {
	ID uint `gorm:"primarykey"`
	*Graph
}

func (GraphEntity) TableName() string {
	return "graph"
}

type NodeEntity struct {
	ID uint `gorm:"primarykey"`
	*Node
	RefGraphId uint `gorm:"ref_graph_id"`
}

func (NodeEntity) TableName() string {
	return "node"
}

type EdgeEntity struct {
	ID uint `gorm:"primarykey"`
	*Edge
	RefFromNodeId uint `gorm:"ref_from_node_id"`
	RefToNodeId   uint `gorm:"ref_to_node_id"`
	RefGraphId    uint `gorm:"ref_graph_id"`
}

func (EdgeEntity) TableName() string {
	return "edge"
}

func (g *GraphEntity) create() error {

	tx := db.Begin()
	if tx.Error != nil {
		logger.Errorf("err when starting tx %s", tx.Error)
		return tx.Error
	}

	// Insert graph
	if err := tx.Create(g).Error; err != nil {
		tx.Rollback()
		logger.Errorf("err when saving graph %s", err)
		return err
	}

	// Map to store node XML IDs to DB IDs
	nodeIDMap := make(map[string]uint)

	// Insert nodes
	for _, node := range g.Nodes {
		n := &NodeEntity{Node: node}
		n.RefGraphId = g.ID
		if err := tx.Create(n).Error; err != nil {
			tx.Rollback()
			logger.Errorf("err when saving nodes %s", err)
			return err
		}
		nodeIDMap[n.Id] = n.RefGraphId
	}

	// Insert edges with mapped node IDs
	for _, edge := range g.Edges {
		e := &EdgeEntity{Edge: edge}
		e.RefFromNodeId = nodeIDMap[e.From]
		e.RefToNodeId = nodeIDMap[e.To]
		e.RefGraphId = g.ID
		if err := tx.Create(e).Error; err != nil {
			tx.Rollback()
			logger.Errorf("err when saving edges %s", err)
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		logger.Errorf("err when commiting tx %s", err)
		return err
	}
	return nil
}

func findLatestGraph() (*Graph, error) {
	ge := &GraphEntity{}
	err := db.Last(ge).Error
	if err != nil {
		return nil, err
	}

	var nodes []*NodeEntity
	err = db.Find(&nodes, NodeEntity{RefGraphId: ge.ID}).Error
	if err != nil {
		return nil, err
	}
	nodeList := make([]*Node, len(nodes))
	for i, n := range nodes {
		node := n.Node
		nodeList[i] = node
	}
	ge.Nodes = nodeList

	var edges []*EdgeEntity
	err = db.Find(&edges, EdgeEntity{RefGraphId: ge.ID}).Error
	if err != nil {
		return nil, err
	}
	edgeList := make([]*Edge, len(edges))
	for i, e := range edges {
		edge := e.Edge
		edgeList[i] = edge
	}
	ge.Edges = edgeList

	return ge.Graph, nil

}
