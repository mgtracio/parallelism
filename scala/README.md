# Parallel Batch Processing in Scala for Challenge Benchmark 

### Steps to run locally:  
Configure Main Class en built.sbt:  
  
```  
mainClass in (Compile, run) := Some("com.latamautos.challenge.benchmark.MainImages")

o

mainClass in (Compile, run) := Some("com.latamautos.challenge.benchmark.MainNumbers")
```
 
Run the following command:  
  
```  
For Images:
    sbt run -mem 2500

For Numbers:
    sbt run -mem 2500
```

