
const { parentPort, workerData, isMainThread } = require("worker_threads");

/**
 * Parallel Batch Processing in Node (Worker Script)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
function workerImages(numbers) {
  const sum = numbers.reduce((sum, current) => sum + current, 0);
  console.log("+++ WORKER FINISHED (sum: %d)\n", sum);
  return sum;
}

if (!isMainThread) { // verify call from worker thread
  if (!Array.isArray(workerData)) throw new Error("workerData must be an array of string"); // throw an error from worker
  parentPort.postMessage(workerImages(workerData));
}