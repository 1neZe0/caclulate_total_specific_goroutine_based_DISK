package DiskUsageCalculate

/*
#include <stdio.h>
#include <sys/resource.h>
#include <sys/time.h>
#include <sys/statvfs.h>
#include <sys/sysinfo.h>

typedef struct {
  struct rusage start_usage;
  struct rusage end_usage;
} DiskUsageCalculate;

void start(DiskUsageCalculate *du) {
  getrusage(RUSAGE_SELF, &du->start_usage);
}

void stop(DiskUsageCalculate *du) {
  getrusage(RUSAGE_SELF, &du->end_usage);
}

long calculate_usage(DiskUsageCalculate *du) {
  long read_bytes = du->end_usage.ru_inblock - du->start_usage.ru_inblock;
  long write_bytes = du->end_usage.ru_oublock - du->start_usage.ru_oublock;
  return read_bytes + write_bytes;
}
double calculate_percentage_of_available_disk_memory_used(DiskUsageCalculate *du) {
  struct statvfs stat;
  if (statvfs(".", &stat) == 0) {
    long long total_available_bytes = (long long) stat.f_bavail * stat.f_frsize;
    long usage_bytes = calculate_usage(du);
    return (double) usage_bytes / total_available_bytes;
  }
  return -1;
}

double calculate_percentage_of_total_disk_memory_used(DiskUsageCalculate *du) {
  struct statvfs stat;
  if (statvfs(".", &stat) == 0) {
    long long total_memory_bytes = (long long) stat.f_blocks * stat.f_frsize;
    long usage_bytes = calculate_usage(du);
    return (double) usage_bytes / total_memory_bytes;
  }
  return -1;
}
*/
import "C"

import (
	"unsafe"
)

type DiskUsageCalculate struct {
	startUsage C.struct_rusage
	endUsage   C.struct_rusage
}

func DiskUsageProfilier() DiskUsageCalculate {
	// initialize fields and return the initialized instance
	return DiskUsageCalculate{
		// field values go here
	}
}
func (du *DiskUsageCalculate) Start() {
	C.start((*C.DiskUsageCalculate)(unsafe.Pointer(du)))
}

func (du *DiskUsageCalculate) Stop() {
	C.stop((*C.DiskUsageCalculate)(unsafe.Pointer(du)))
}

func (du *DiskUsageCalculate) CalculateUsage() int64 {
	return int64(C.calculate_usage((*C.DiskUsageCalculate)(unsafe.Pointer(du))))
}

func (du *DiskUsageCalculate) CalculatePecentageOfAvailableDiskMemoryUsed() float64 {
	return float64(C.calculate_percentage_of_available_disk_memory_used((*C.DiskUsageCalculate)(unsafe.Pointer(du))))
}

func (du *DiskUsageCalculate) CalculatePecentageOfTotalDiskMemoryUsed() float64 {
	return float64(C.calculate_percentage_of_total_disk_memory_used((*C.DiskUsageCalculate)(unsafe.Pointer(du))))
}

//func main() {
//	var du DiskUsageCalculate
//
//	du.Start()
//	// Do some disk I/O here
//	du.Stop()
//
//	usage := du.CalculateUsage()
//	fmt.Printf("Disk usage: %d bytes\n", usage)
//}
