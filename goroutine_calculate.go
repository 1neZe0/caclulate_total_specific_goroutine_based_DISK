package main

import (
	"awrpoj/diskusageprofilier"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	profiler := DiskUsageCalculate.DiskUsageProfilier()
	profiler.Start()

	// Create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Start the goroutine
		f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		for i := 0; i < 100000000; i++ {
			_, _ = f.WriteString("gdfgdfg")
		}
	}()

	// Wait for the goroutine to finish
	wg.Wait()

	profiler.Stop()
	usage := profiler.CalculateUsage()
	fmt.Printf("Disk usage: %d bytes\n", usage)

	po_of_av := profiler.CalculatePecentageOfAvailableDiskMemoryUsed()
	fmt.Printf("Disk usage percentage of available disk memory: %.10f%%\n", po_of_av)
	formattedUsage := strconv.FormatFloat(po_of_av, 'f', 10, 64)
	parsedUsage, _ := strconv.ParseFloat(formattedUsage, 64)
	maximum_number_of_goroutines := 100 / parsedUsage
	fmt.Printf("Maximum number of goroutines based on available disk memory: %.f\n", maximum_number_of_goroutines)

	po_of_to := profiler.CalculatePecentageOfTotalDiskMemoryUsed()
	fmt.Printf("Disk usage percentage of total disk memory: %.10f%%\n", po_of_to)
	formattedUsage = strconv.FormatFloat(po_of_to, 'f', 10, 64)
	parsedUsage, _ = strconv.ParseFloat(formattedUsage, 64)
	maximum_number_of_goroutines = 100 / parsedUsage
	fmt.Printf("Maximum number of goroutines based on total disk memory: %.f\n", maximum_number_of_goroutines)

}
