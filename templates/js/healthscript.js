let getvalue = document.getElementById('symptom');
let drug = document.getElementById('displaydrugs');

getvalue.addEventListener('input', show);

// Shows the recommended drug in the displaydrugs div
function show() {
    if (getvalue.value === "Bloody droppings/diarrhoea") {
        drug.innerHTML = "amprolium, sulfadimethoxine and trimethoprim";
    } else if (getvalue.value === "Swollen face/eyes, Sticky discharge from the eyes and nostrils") {
        drug.innerHTML = "Isolation and treatment with water soluble antibacterials can be used mainly erythromycin, tetracycline and sulfa drugs.";
    } else if (getvalue.value === "Loss of appetites but increased thirst, gasping or convulsions") {
        drug.innerHTML = "Proper hygiene all round. Improved ventilation";
    } else if (getvalue.value === "Wart-like growths or sores. Slow growth and canker in the mouth") {
        drug.innerHTML = "Fowl Pox Vaccine";
    } else if (getvalue.value === "Dull eyes, head and neck tremors, appetite loss and eventual paralysis") {
        drug.innerHTML = "Avian Encephalomyelitis Vaccine";
    } else {

    }
}