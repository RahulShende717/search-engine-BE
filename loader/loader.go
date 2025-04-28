package loader

import (
	"log"

	"search-eng/search"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func LoadParquetFile(filePath string) []search.Record {
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		log.Fatalf("Can't open file: %v", err)
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(search.Record), 4)
	if err != nil {
		log.Fatalf("Can't create parquet reader: %v", err)
	}
	defer pr.ReadStop()

	//  Add this block to print the Parquet schema
	// fmt.Println("Parquet Schema:")
	// for i, schema := range pr.SchemaHandler.SchemaElements {
	// 	fmt.Printf("[%d] Name: %s, Type: %v, RepetitionType: %v\n", i, schema.GetName(), schema.Type, schema.RepetitionType)
	// }

	num := int(pr.GetNumRows())
	records := make([]search.Record, num)

	if err = pr.Read(&records); err != nil {
		log.Fatalf("Read error: %v", err)
	}

	return records
}
