const lineChartOptions = {
  chart: {
    type: 'bar',
    fontFamily: 'Inter, sans-serif',
    height: 550,
    toolbar: {
      show: true,
    },
    animations: {
      enabled: false,
    },
  },
  plotOptions: {
    bar: {
      horizontal: false,
      columnWidth: '80%',
      endingShape: 'rounded',
    },
  },
  dataLabels: {
    enabled: false,
  },
  stroke: {
    show: true,
    width: 2,
  },
  xaxis: {
    labels: {
      style: {
        fontSize: '14px',
      },
    },
    tooltip: {
      enabled: true,
      offsetY: 25,
      formatter: function (val, opts) {
        switch (val) {
          case 'Event':
            return 'Log a message with no context';
          case 'EventFmt':
            return 'Log a message with string-formatted variables';
          case 'EventCtx':
            return 'Log a message with contexual fields';
          case 'EventCtxWeak':
            return 'Log a message with loosely-typed contexual fields';
          case 'EventAccumulatedCtx':
            return 'Log a message with accumulated contexual fields';
          case 'Disabled':
            return 'Log a message at a disabled level';
          case 'DisabledFmt':
            return 'Log at a disabled level with string-formatted variables';
          case 'DisabledCtx':
            return 'Log at a disabled level with contextual fields';
          case 'DisabledCtxWeak':
            return 'Log at a disabled level with loosely-typed contextual fields';
          case 'DisabledAccumulatedCtx':
            return 'Log at a disabled level with accumulated contextual fields';
          default:
            return val;
        }
      },
    },
  },
  yaxis: [
    {
      labels: {
        style: {
          fontSize: '14px',
        },
      },
      decimalsInFloat: 0,
      title: {
        style: {
          fontSize: '13px',
        },
      },
    },
  ],
  fill: {
    opacity: 1,
  },
  title: {
    text: 'Execution time',
    margin: 20,
    style: {
      fontSize: '24px',
    },
  },
  subtitle: {
    text: '',
    margin: 30,
    style: {
      fontSize: '16px',
    },
  },
  legend: {
    fontSize: '16px',
    position: 'bottom',
    itemMargin: {
      horizontal: 15,
      vertical: 25,
    },
  },
  colors: [
    '#008FFB',
    '#00E396',
    '#FF9800',
    '#E91E63',
    '#775DD0',
    '#5653FE',
    '#FD6A6A',
    '#66DA26',
    '#546E7A',
    '#2E294E',
  ],
};

export default function (obj) {
  lineChartOptions.series = obj.series;
  lineChartOptions.xaxis.categories = obj.categories;
  lineChartOptions.yaxis[0].title.text = obj.yaxis;
  lineChartOptions.title.text = obj.title;
  lineChartOptions.subtitle.text = obj.subtitle || 'A lower score is better';

  return lineChartOptions;
}
