package com.latamautos.challenge.benchmark

import java.util.concurrent.Executors
import com.latamautos.challenge.benchmark.Config._
import scala.concurrent._

/**
 * Parallel Batch Processing in Scala (MainNumbers)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
object MainNumbers {
  val executor = Executors.newWorkStealingPool()
  implicit val executionContext: ExecutionContextExecutor = ExecutionContext.fromExecutor(executor)

  def main(args: Array[String]) = {
    def worker(numbers: List[Int]) = {
      val sum = numbers.fold(0){ (a, b) => a + b }
      printf("+++ WORKER FINISHED (sum: %d of %d elements)\n", sum)
      sum
    }

    val start = System.currentTimeMillis
    val groupWorkers: Future[List[Int]] = Future.sequence(numbers.grouped((numbers.size / WorkersForNumbers.toFloat).ceil.toInt).map(x => Future(worker(x))).toList)
    groupWorkers.onComplete(result => {
      println(s"::: INTENSIVE JOB SUCCESSFULLY COMPLETED -> Workers: ${result.getOrElse(Nil).size} -> Elapsed time: ${System.currentTimeMillis - start} ms -> Total Sum: ${result.getOrElse(Nil).fold(0){ (a, b) => a + b }}")
      executor.shutdownNow()
    }) // Wait for Futures Group to finish

  }
}