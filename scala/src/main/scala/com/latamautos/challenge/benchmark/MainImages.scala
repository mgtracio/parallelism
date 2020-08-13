package com.latamautos.challenge.benchmark

import com.latamautos.challenge.benchmark.Config._
import java.io.{File, FileInputStream}
import com.sksamuel.scrimage.ImmutableImage
import com.sksamuel.scrimage.nio.JpegWriter
import scala.concurrent._
import java.util.concurrent.Executors

/**
 * Parallel Batch Processing in Scala (MainImages)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
object MainImages {
  implicit val ec = ExecutionContext.fromExecutorService(Executors.newFixedThreadPool(images.size))

  def main(args: Array[String]) = {
    def worker(images: List[File]) = {
      images.foreach(i => ImmutableImage.loader().fromStream(new FileInputStream(i)).fit(X, Y).output(new JpegWriter().withCompression(Quality), Config.imageWriter(i)))
      printf("+++ WORKER FINISHED (processed images: %d)\n", images.size)
    }

    val start = System.currentTimeMillis
    val groupWorkers =  Future.sequence(images.grouped((images.size / WorkersForImages.toFloat).ceil.toInt).map(x => Future(worker(x))))
    groupWorkers.onComplete(result => printf(s"::: INTENSIVE JOB SUCCESSFULLY COMPLETED -> Workers: ${result.getOrElse(Nil).size} -> Elapsed time: ${System.currentTimeMillis - start} ms")) // Wait for Futures Group to finish
  }
}