const addButton = document.getElementById('additem');
const addExpButton = document.getElementById('additemexp');
//let pickedAccDate = document.getElementById('acdate');
//let accDate = document.getElementById('accdate');
let summary = document.getElementById('display');
let salesItemInput = document.getElementById('birdsold'); // Gets the sales item
let salesPriceInput = document.getElementById('sellprice'); // Gets the sales price
let salesQtyInput = document.getElementById('qtysold'); // Gets the sales qty
let expItemInput = document.getElementById('exp'); // Gets the expense item
let expPriceInput = document.getElementById('expprice'); // Gets the expense price
let expQtyInput = document.getElementById('qtyexp'); // Gets the expense qty

//accDate.addEventListener(onchange, replace);
addButton.addEventListener('click', render);
addExpButton.addEventListener('click', renderExpenses);
/*
function replace() {
    pickedAccDate.innerHTML = accDate.value;
}*/

// Shows the inputted sales in the summary div
function render() {
    executeSalesOnce();
    summary.innerHTML += `${salesItemInput.value} &nbsp &nbsp &nbsp &nbsp ${salesQtyInput.value} &nbsp &nbsp &nbsp $${salesPriceInput.value * salesQtyInput.value}<br>`;
    //addSales();
    salesItemInput.value = '';
    salesQtyInput.value = '';
    salesPriceInput.value = '';

}

// Prints Sales once
function executeSalesOnce() {
    summary.innerHTML += 'Sales<br>';
    executeSalesOnce = function () { };
}
/*
function addSales() {
    totalSales = salesPriceInput.value * salesQtyInput.value;
    totalSales += 0;
    summary.innerHTML += `Total &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp${totalSales}`
}*/

// Shows the inputted expenses in the summary div
function renderExpenses() {
    executeExpOnce();
    summary.innerHTML += `${expItemInput.value} &nbsp &nbsp &nbsp &nbsp ${expQtyInput.value} &nbsp &nbsp &nbsp $${expPriceInput.value * expQtyInput.value}<br>`;
    //addExpenses();
    expItemInput.value = '';
    expQtyInput.value = '';
    expPriceInput.value = '';
}

// Prints Expense once
function executeExpOnce() {
    summary.innerHTML += 'Expenses<br>';
    executeExpOnce = function () { };
}
