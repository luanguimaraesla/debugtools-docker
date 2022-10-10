#! /bin/bash

# This program uses fio to generate useful insights about disk
# usage by performing I/O testing based on the fio.ini configuration
# file. Then it generates charts summarizing the data collected
# in .png files using fio-plot.
#
# Read more:
# https://fio.readthedocs.io/en/latest/fio_doc.html
# https://github.com/louwrentius/fio-plot

echo "Cleaning old results"
rm -rf ./results

echo "Creating testing folder: ./results"
mkdir ./results

echo "Running FIO"
fio --output=results.json --output-format=json fio.ini

echo "Running FIO plot for IOPS"
fio-plot \
	--input-directory="./results" \
	--output-filename="./results.iops.png" \
	--title="Disk R/W reliability benchmark - IOPS" \
	--source="luan@timescale.com" \
	--loggraph \
	--numjobs=4 \
	--iodepth=64 \
	--rw="randrw" \
	--type=iops \
	--xlabel-parent=0

echo "Running FIO plot for latency"
fio-plot \
	--input-directory="./results" \
	--output-filename="./results.lat.png" \
	--title="Disk R/W reliability benchmark - Latency" \
	--source="luan@timescale.com" \
	--loggraph \
	--numjobs=4 \
	--iodepth=64 \
	--rw="randrw" \
	--type=lat \
	--xlabel-parent=0

echo "Finishing"
