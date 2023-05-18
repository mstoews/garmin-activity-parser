# garmin-activity-parser

## Running on the Command Line
To run the program that exacts the data from the tcx file, run 

```bash
./main -filename=./xml/activity_3873231072.tcx
```

The file location must be entered relative to the program. If the activity is in the same directory 

```bash
./main -filename=./activity_3873231072.tcx
```

The file is comma separated so if you want to write to a file, just use the standard OS pipe symbol

```bash
./main -filename=./activity_3873231072.tcx > stats.csv
```

The comma separated file can then be added to Excel or your analysis tools.  

## Data Cleanup
The data will need clean up as there are pauses in the data capture process on the garmin side. 




