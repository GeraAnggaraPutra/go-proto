package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	pb "go-proto/pb"

)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 1000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Nike Air Jordan",
				Price: 1200.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("marshall error", err)
	}

	// compact binary wire format
	fmt.Println(data)

	p := &pb.Products{}
	if err = proto.Unmarshal(data, p); err != nil {
		log.Fatal("unmarshall error", err)
	}

	for _, product := range p.GetData() {
		fmt.Println(product.GetId())
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory().GetId())
		fmt.Println(product.GetCategory().GetName())
	}

	fmt.Println(p)
}
