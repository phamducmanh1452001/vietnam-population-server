var apiGenderAdmin, apiAgeAdmin, apiReligionAdmin, codeAdmin;
if(localStorage.getItem('level')==0){
    apiGenderAdmin = "https://www.phorifai.xyz/api/gender-chart";
    apiReligionAdmin = "https://www.phorifai.xyz/api/religion-chart";
    apiAgeAdmin = "https://www.phorifai.xyz/api/age-chart";
}
else if(localStorage.getItem('level')==1){
    codeAdmin = localStorage.getItem('user');
    apiGenderAdmin = "https://www.phorifai.xyz/api/gender-chart?code="+codeAdmin;
    apiReligionAdmin = "https://www.phorifai.xyz/api/religion-chart?code="+codeAdmin;
    apiAgeAdmin = "https://www.phorifai.xyz/api/age-chart?code="+codeAdmin;    
}
else if(localStorage.getItem('level')==2){
    codeAdmin = localStorage.getItem('user');
    apiGenderAdmin = "https://www.phorifai.xyz/api/gender-chart?code="+codeAdmin;
    apiReligionAdmin = "https://www.phorifai.xyz/api/religion-chart?code="+codeAdmin;
    apiAgeAdmin = "https://www.phorifai.xyz/api/age-chart?code="+codeAdmin;
}
else{
    codeAdmin = localStorage.getItem('user');
    apiGenderAdmin = "https://www.phorifai.xyz/api/gender-chart?code="+codeAdmin;
    apiReligionAdmin = "https://www.phorifai.xyz/api/religion-chart?code="+codeAdmin;
    apiAgeAdmin = "https://www.phorifai.xyz/api/age-chart?code="+codeAdmin;
}

   // TÍnh tỷ lệ giới tính
//    let chartEleAdmin = '<canvas id="chartSex-admin" ></canvas>';
//    $('#chartSexX').append(chartEleAdmin);
    fetch(apiGenderAdmin)
    .then(response=>response.json())
    .then(dataGender=>{
        // hiển thị biểu đồ theo giới tính
        let xValues = ["Nam", "Nữ"];
        let yValues = [dataGender.nam, dataGender.nu];
        let barColors = [
          "#17a2b8",
          "#6f42c1"
        ];

        // Biểu đồ tỷ lệ giới tính
        new Chart("chartSex-admin", {
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
              text: "Biểu đồ dân số theo giới tính",
              position: "bottom",
              fontSize: 60
            },
            legend:{
              display: true,
              labels: {
                fontSize: 45               
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

        //Biều đồ dân số theo tôn giáo
    fetch(apiReligionAdmin)
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

        new Chart("chartEthnic-admin", {
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
                    text: "Biểu đồ cơ cấu theo tôn giáo năm 2020",
                    position: "bottom",
                    fontSize: 60
                },
                legend:{
                    display: true,
                    labels:{
                        fontSize: 45
                    }
                }
            }
        });
    })
// Biểu đồ theo độ tuổi
    fetch(apiAgeAdmin)
    .then(response=>response.json())
    .then(dataAge=>{
        var xValues = ["0-10", "10-20", "20-30", "30-40", "40-50","50-60","60-70","70-80","80-90","90 trở lên"];
        var yValues = [];
        Chart.defaults.global.defaultFontSize = 20;
        for(let key in dataAge){
            yValues.push(dataAge[key]);
        }
        new Chart("chartAge-admin", {
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
                    text: "Biểu đồ phân chia dân số theo độ tuổi",
                    position: 'bottom',
                    fontSize: 20
                }
            }
        });
    })