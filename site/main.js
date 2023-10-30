import ApexCharts from 'apexcharts';
import chart from './chart.js';
import data from '../bench.json';

let categories = [];
let libraries = [];

const series = {
  executionTime: [],
  executionTimeDisabled: [],
  memoryUsage: [],
  memoryUsageDisabled: [],
  totalRuns: [],
  totalRunsDisabled: [],
  allocations: [],
  allocationsDisabled: [],
};

function pushToCharts(prop, lib, value) {
  console.log(prop, lib, value);
  const i = series[prop].findIndex((obj) => obj.name === lib);
  if (i === -1) {
    series[prop].push({
      name: lib,
      data: [value],
    });
  } else {
    series[prop][i].data.push(value);
  }
}

const benchmarks = data[0].Suites[0].Benchmarks;

benchmarks.forEach((item) => {
  const benchName = item.Name.split('/')[0].split('Benchmark')[1];
  const library = item.Name.split('/')[1].split('-')[0];
  categories.push(benchName);
  libraries.push(library);

  console.log(libraries);

  if (!benchName.includes('Disabled')) {
    pushToCharts('executionTime', library, item.NsPerOp);
    pushToCharts('memoryUsage', library, item.Mem.BytesPerOp);
    pushToCharts('totalRuns', library, item.Runs);
    pushToCharts('allocations', library, item.Mem.AllocsPerOp);
  } else {
    pushToCharts('executionTimeDisabled', library, item.NsPerOp);
    pushToCharts('memoryUsageDisabled', library, item.Mem.BytesPerOp);
    pushToCharts('totalRunsDisabled', library, item.Runs);
    pushToCharts('allocationsDisabled', library, item.Mem.AllocsPerOp);
  }
});

const uniqueCategories = Array.from(new Set(categories));

const enabledCategories = uniqueCategories.filter(
  (e) => !e.includes('Disabled')
);
const disabledCategories = uniqueCategories.filter((e) =>
  e.includes('Disabled')
);

const executionTimeChart = new ApexCharts(
  document.querySelector('#js-nano-chart'),
  chart({
    series: series.executionTime,
    categories: enabledCategories,
    title: 'Execution time',
    subtitle: 'Average execution time per logged event (lower is better)',
    yaxis: 'nanoseconds',
  })
);
executionTimeChart.render();

const executionTimeDisabledChart = new ApexCharts(
  document.querySelector('#js-nano-chart-disabled'),
  chart({
    series: series.executionTimeDisabled,
    categories: disabledCategories,
    title: 'Execution time (disabled)',
    subtitle: 'Average execution time per disabled log (lower is better)',
    yaxis: 'nanoseconds',
  })
);
executionTimeDisabledChart.render();

const allocsChart = new ApexCharts(
  document.querySelector('#js-allocs-chart'),
  chart({
    series: series.allocations,
    categories: enabledCategories,
    title: 'Heap allocations',
    subtitle: 'Incurred allocations per logged event (lower is better)',
    yaxis: 'allocations',
  })
);
allocsChart.render();

const allocsDisabledChart = new ApexCharts(
  document.querySelector('#js-allocs-chart-disabled'),
  chart({
    series: series.allocationsDisabled,
    categories: disabledCategories,
    title: 'Heap allocations (disabled)',
    subtitle: 'Incurred allocations per disabled event (lower is better)',
    yaxis: 'allocations',
  })
);
allocsDisabledChart.render();

const memoryUsageChart = new ApexCharts(
  document.querySelector('#js-bytes-chart'),
  chart({
    series: series.memoryUsage,
    categories: enabledCategories,
    title: 'Memory usage',
    subtitle: 'Bytes allocated per logged event (lower is better)',
    yaxis: 'bytes',
  })
);
memoryUsageChart.render();

const memoryUsageDisabledChart = new ApexCharts(
  document.querySelector('#js-bytes-chart-disabled'),
  chart({
    series: series.memoryUsageDisabled,
    categories: disabledCategories,
    title: 'Memory usage (disabled)',
    subtitle: 'Bytes allocated per disabled event (lower is better)',
    yaxis: 'bytes',
  })
);
memoryUsageDisabledChart.render();

const totalOpsChart = new ApexCharts(
  document.querySelector('#js-run-chart'),
  chart({
    series: series.totalRuns,
    categories: enabledCategories,
    title: 'Total events',
    subtitle: 'Number of events logged (higher is better)',
    yaxis: 'iterations',
  })
);
totalOpsChart.render();

const totalOpsDisabledChart = new ApexCharts(
  document.querySelector('#js-run-chart-disabled'),
  chart({
    series: series.totalRunsDisabled,
    categories: disabledCategories,
    title: 'Total events (disabled)',
    subtitle: 'Number of disabled events skipped (higher is better)',
    yaxis: 'iterations',
  })
);
totalOpsDisabledChart.render();
