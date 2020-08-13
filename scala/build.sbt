name := "ChallengeBenchmarkScala"

version := "0.1"

scalaVersion := "2.13.3"

libraryDependencies += "com.sksamuel.scrimage" % "scrimage-core" % "4.0.5"

//mainClass in (Compile, run) := Some("com.latamautos.challenge.benchmark.MainImages")
mainClass in (Compile, run) := Some("com.latamautos.challenge.benchmark.MainNumbers")