//Biểu đồ 2010-2020
var r=1;
// Năm
var year = ["2010", "2011", "2012", "2013", "2014", "2015","2016", "2017", "2018", "2019","2020"];
// Giá trị dân số qua các năm
var populationYear = [7666, 8189, 8990, 9000, 11000, 11034, 12955, 13949, 14044, 15024, 16700];
var barColors = "#20c997";

function chartYear(x,y,bar,id){
  new Chart(document.getElementById(id), {
    // new Chart(document.getElementById("mixed-chart"), {
      type: 'bar',
      data: {
        labels: x,
        datasets: [{
            type: "line",
            borderColor: "#8e5ea2",
            data: y,
            fill: false
          }, {
            type: "bar",
            backgroundColor: bar,
            backgroundColorHover: "#3e95cd",
            data: y
          }
        ]
      },
      options: {
          title: {
              display: true,
              text: 'Biểu đồ dân số Việt Nam từ năm 2010-2020',
              position: 'bottom',
              fontSize: 16,
              color:"#3e95cd"
          },
          legend: { display: true }                
      }
  });  
}
chartYear(year,populationYear,barColors,'mixed-chart');
// Biểu đồ dân tộc
// Tên dân tộc
fetch("https://www.phorifai.xyz/api/religion-chart")
.then(response=>response.json())
.then(dataReligion=>{
  var xValues = ["Cao Đài", "Công Giáo", "Hòa Hảo", "Hồi Giáo", "Khác", "Không", "Phật giáo","Tin lành","Tôn giáo dân gian"];
  // Dân số qua các dân tộc
  var yValues = [dataReligion.cao_dai, dataReligion.cong_giao , dataReligion.hoa_hao, dataReligion.hoi_giao, dataReligion.khac, dataReligion.khong,dataReligion.phat_giao,dataReligion.tin_lanh,dataReligion.ton_giao_dan_gian];
  var barColors = [
    "#5d767e",
    "#00aba9",
    "#2b5797",
    "#e8c3b9",
    "#1e7145",
    "#ffc107",
    "#fd7e14",
    "#6c757d",
    "#b91d47"
  ];

  new Chart("myChart1", {
    type: "pie",
    data: {
      labels: xValues,
      datasets: [{
        backgroundColor: barColors,
        data: yValues
      }]
    },
    options: {
      title: {
        display: true,
        text: "Biểu đồ cơ cấu theo tôn giáo của nước ta năm 2020",
        position: "bottom",
        fontSize: 20
      },
      legend:{
        display: true,
      }
    }
  });
})

fetch("https://www.phorifai.xyz/api/gender-chart")
.then(response=>response.json())
.then(dataGender=>{
  let xValues = ["Nam", "Nữ"];
  let yValues = [dataGender.nam, dataGender.nu];
  let barColors = [
    "#17a2b8",
    "#6f42c1"
  ];
  // Biểu đồ tỷ lệ giới tính
  new Chart("myChart2", {
    type: "pie",
    data: {
      labels: xValues,
      datasets: [{
        backgroundColor: barColors,
        data: yValues
      }]
    },
    options: {
      title: {
        display: true,
        text: "Biểu đồ dân số theo giới tính của nước ta",
        position: "bottom",
        fontSize: 20
      },
      legend:{
        display: true,
        labels: {
        }
      },
      layout:{
        padding:{
          top: 20
        }      
      }
    }
  });
})
fetch("https://www.phorifai.xyz/api/age-chart")
.then(response=>response.json())
.then(dataAge=>{
  // Biểu đồ giới tính
    var xValues = ["0-10 tuổi", "10-20 tuổi", "20-30 tuổi", "30-40 tuổi", "40-50 tuổi","50-60 tuổi","60-70 tuổi","70-80 tuổi","80-90 tuổi","90 tuổi trở lên"];
    var yValues = [];
  for(let key in dataAge){
    yValues.push(dataAge[key]);
  }
  new Chart("myChart3", {
    type: "bar",
    data: {
      labels: xValues,  
      datasets: [{
        label: 'Dân số',
        backgroundColor: "#2b5797",
        data: yValues
      }]
      },
    options: {
      legend: {
        display: true,
      },
      title: {
        display: true,
        text: "Biểu đồ phân chia dân số theo độ tuổi 2020",
        position: 'bottom',
        fontSize: 20
      }
    }
  });
})