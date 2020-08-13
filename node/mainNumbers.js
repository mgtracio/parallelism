const { Worker } = require('worker_threads');
const config = require('./config');

/**
 * Parallel Batch Processing in Node (mainNumbers)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
async function mainNumbers() {
    function worker(numbers) {
        return new Promise((resolve, reject) => {
            const worker = new Worker(config.workerFileNumbers, { workerData: numbers });
            worker.on(config.workerEventMessage, resolve);
            worker.on(config.workerEventError, reject);
        });
    }

    const start = Date.now();
    const numbers = config.numbers
    const segmentsPerWorker = Math.ceil(numbers.length / config.workersForNumbers);
    const batches = Array(config.workersForNumbers).fill().map((_, index) => {
        let images;
        if (index === 0) {
            images = numbers.slice(0, segmentsPerWorker);
        } else if (index === config.workersForNumbers - 1) {
            images = numbers.slice(segmentsPerWorker * index);
        } else {
            images = numbers.slice(segmentsPerWorker * index, segmentsPerWorker * (index + 1))
        }
        return worker(images)
    });
    Promise.all(batches).then(result => console.log(`::: INTENSIVE JOB SUCCESSFULLY COMPLETED -> Workers: ${config.workersForNumbers} -> Elapsed time: ${Date.now() - start} ms -> Total Sum: ${result.reduce((accumulator, currentValue) => accumulator + currentValue)}`)); // Wait for Workers and Promises Group to finish
}

mainNumbers();