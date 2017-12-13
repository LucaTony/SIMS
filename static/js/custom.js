/*Ajax live search*/
function showResult(str) {
    if (str.length==0) {
        document.getElementById("liveresults").innerHTML="";
        return;
    }
    if (window.XMLHttpRequest) {
        // code for IE7+, Firefox, Chrome, Opera, Safari
        xmlhttp=new XMLHttpRequest();
    } else {  // code for IE6, IE5
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange=function() {
        if (this.readyState==4 && this.status==200) {
            document.getElementById("liveresults").innerHTML=this.responseText;
        }
    }
    xmlhttp.open("GET","ajax/"+str,true);
    xmlhttp.send();
}


/*Quiz Radio with Accordion compatibility*/
$(document).ready(function () {
    $("[id^='r_']").on('click', function(){
        $(this).parent().find('a').trigger('click')
    })

});


(function(){
    new Clipboard('#copy-button');
})();

function changeCard(quest, np) {
    var maxquest = 10;
    var parentpro = document.getElementById("pro");
    var childpro = document.getElementById("probar");
    var cards = document.getElementsByClassName("calc-card");

    //Check if at least one radio has been checked
    //ignore if you want to start the quiz or go back
    if ($('input[name=Question'+quest+']:checked').length > 0  || quest == 0 || np == 'prev') {
        if (np == "next") quest++;
        else if (np == "prev") quest--;

        for (var i = 0; i <cards.length; i++ ){
            cards[i].style.display = "none";
        }
        var newCard = document.getElementById("Question"+quest);
        newCard.style.display = "block"; 
        if(quest == 11){
            parentpro.style.display = "none";
        } else {
            parentpro.style.display = "block";
            w = quest / maxquest * 100;
            childpro.innerText = quest + "/" + maxquest; 
            childpro.style.width = w + "%";
        }
    } else  alert("Vänligen välj alternativ");

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
    var one = document.getElementById("fact" + getRand());
    //hides all generated facts and shows one specific.
    for (var i = 0; i < facts.length; i++) {
        facts[i].style.display = "none";

    }
    one.style.display = "block";
    timer();

}
function getRand(){
    rand = Math.floor(Math.random() * 25) + 1;
    console.log(rand);
    return rand;
}
//The time each fact will be showed
function timer() {
    setInterval(function () {
        for(var i = 1; i<26; i++){
            var first = document.getElementById("fact" + i);
            first.style.display = "none";
        }
        var one = document.getElementById("fact" + getRand());
        one.style.display = "block";

    }, 15000);
}
