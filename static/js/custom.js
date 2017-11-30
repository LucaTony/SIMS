
function changeCard(quest, np) {
    var maxquest = 10;
    var parentpro = document.getElementById("pro");
    var childpro = document.getElementById("probar");
    var cards = document.getElementsByClassName("calc-card");
    if (np == "next") quest++;
    else if (np == "prev") quest--;

    for (var i = 0; i < cards.length; i++) {
        cards[i].style.display = "none";
    }
    var newCard = document.getElementById("Question" + quest);
    newCard.style.display = "block";
    if (quest == 11) {
        parentpro.style.display = "none";
    } else {
        parentpro.style.display = "block";
        w = quest / maxquest * 100;
        childpro.innerText = quest + "/" + maxquest;
        childpro.style.width = w + "%";
    }

}

function sticky_close() {
    var sticky = document.getElementById("sticky");
    sticky.style.display = "none";
}


function change(card) {

    var infoCard = document.getElementsByClassName("card bg-light");
    var newCard = document.getElementById(card);

    for (var i = 0; i < infoCard.length; i++) {
        infoCard[i].style.display = "none";
    }
    newCard.style.display = "block";
}


$(document).ready(function () {
    $("#cf_controls").on('click', 'span', function () {
        $("#cf img").removeClass("opaque");

        var newImage = $(this).index();

        $("#cf img").eq(newImage).addClass("opaque");

        $("#cf_controls span").removeClass("selected");
        $(this).addClass("selected");
    });
});

//Display new fact and hide all other facts
/*after ten seconds hide recent fact and show new fact. */
function changeFact() {
    var facts = document.getElementsByClassName("funFact");
    var one = document.getElementById("fact" + 1);
    //hides all generated facts and shows one specific.
    for (var i = 0; i < facts.length; i++) {
        facts[i].style.display = "none";

    }
    one.style.display = "block";
    timer();

}

//The time each fact will be showed
function timer() {
    setInterval(function () {
        for(var i = 1; i<6; i++){
        var first = document.getElementById("fact" + i);
        first.style.display = "none";
        }
        rand = Math.floor(Math.random() * 5) + 1;
        var one = document.getElementById("fact" + rand);
        one.style.display = "block";

    }, 1000);
}
