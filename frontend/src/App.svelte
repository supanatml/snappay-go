<script>
  import {QRpopup} from '../wailsjs/go/main/App.js'
  import {CheckTransaction} from '../wailsjs/go/main/App.js'
  import {VoucherPopup} from '../wailsjs/go/main/App.js'
  import {PrintLog} from '../wailsjs/go/main/App.js'

  let page = 0
  let qrPromise
  let timeoutID
  let intervalID
  let ref = "0"
  let voucherinput = ""
  let timeout = 0

  //todo: pass ref instead of global var?

  function showChoice() {
    page = 1
    clearInterval(intervalID)
    clearTimeout(timeoutID)
  }
  async function showQR() {
    page = 2
    timeout = 120
    ref = await QRpopup();
	  timeoutID = setTimeout(() => page = 0, timeout*1000);
      intervalID = setInterval(function() {
        timeout--
      },1000)
    return ref
  }
  async function showVoucher() {
    page = 3
    voucherinput = ""
    timeout = 30
    timeoutID = setTimeout(() => page = 0, timeout*1000);
    intervalID = setInterval(function() {
      timeout--
    },1000)
    await VoucherPopup()
  }
  function checkPaid() {
    let paid = CheckTransaction(ref)
    if(paid) {
      clearTimeout(timeoutID)
    }
  }
  function typein(s) {
    var label = document.getElementById("voucherinput")
    x.innerHTML = x.innerHTML + s
  }
</script>

<main>
  {#if page === 0}
    <div class="btn-group">
      <button class="btn start" on:click={showChoice}>Start</button>
    </div>
  {:else if page === 1}
    <button class="btn pay" on:click={() => qrPromise = showQR()}>Pay</button>
    <button class="btn voucher" on:click={showVoucher}>Voucher</button>
  {:else if page === 2}
    {#await qrPromise}
      <h1>Loading...</h1>
    {:then ref}
      <img id="qr" class="qr" src="qr.png?ref={ref}" alt="No QR available"/>
    {/await}
      <button class="btn paid" on:click={checkPaid}>Confirm payment</button>
      <h1 class="timeout">{timeout} seconds</h1>
  {:else if page === 3}
    <div class="voucherpage">
      <button class="btn back" on:click={() => (page = 0)}>Back</button>
      <label class="voucherinput">{voucherinput}</label>
      <div class="numpad">
        <button class="num n0" on:click={() => voucherinput += 0}>0</button>
        <button class="num n1" on:click={() => voucherinput += 1}>1</button>
        <button class="num n2" on:click={() => voucherinput += 2}>2</button>
        <button class="num n3" on:click={() => voucherinput += 3}>3</button>
        <button class="num n4" on:click={() => voucherinput += 4}>4</button>
        <button class="num n5" on:click={() => voucherinput += 5}>5</button>
        <button class="num n6" on:click={() => voucherinput += 6}>6</button>
        <button class="num n7" on:click={() => voucherinput += 7}>7</button>
        <button class="num n8" on:click={() => voucherinput += 8}>8</button>
        <button class="num n9" on:click={() => voucherinput += 9}>9</button>
        <button class="num del" on:click={() => voucherinput = voucherinput.slice(0,-1)}>Del</button>
        <button class="num ok">OK</button>
      </div>
      <h1 class="timeout">{timeout} seconds</h1>
    </div>
  {/if}
</main>

<style>
  #qr {
    display: block;
    height: 450px;
    position: relative;
    left: 38vw;
  }
  .timeout {
    margin-top: 0;
    grid-area: timeout;
  }
  .btn {
    width: 350px;
    height: 120px;
    margin-bottom: 20px;
    font-size: 0;
    border: none;
  }

  .start {
    background:url(./assets/start.png) no-repeat;
    background-size: contain;
    background-position: center;
    margin-top: 260px;
  }
  .pay {
    background:url(./assets/pay.png) no-repeat;
    background-size: contain;
    background-position: center;
    margin-top: 160px;
  }
  .voucherpage {
    display: grid;
    grid-template-areas:
      "back"
      "voucherinput"
      "numpad"
      "timeout";
    gap: 0 0;
    align-items: center;
    justify-items: center;
  }
  .voucher {
    background:url(./assets/voucher.png) no-repeat;
    background-size: contain;
    background-position: center;
    margin-top: 160px;
  }
  .back {
    background:url(./assets/back.png) no-repeat;
    background-size: contain;
    background-position: center;
    width: 200px;
    grid-area: back;
    position: absolute;
    left: 0;
    top: 0;
  }
  .voucherinput {
    grid-area: voucherinput;
    padding: 50px 50px;
    font-size: 4rem;
    margin-top: 100px;
    margin-bottom: 100px;
    height: 4rem;
    width: 450px;
    border: 4px solid white;
    border-radius: 20px;
    overflow: hidden;
  }
  .numpad {
    grid-area: numpad;
    display: grid;
    grid-template-areas:
      "n0 n2 n4 n6 n8 del"
      "n1 n3 n5 n7 n9 ok";
    gap: 0 0;
    align-items: center;
    justify-items: center;
  }
  .num {
    background:url(./assets/blank.png) no-repeat;
    background-size: contain;
    background-position: center;
    font-size: 2rem;
    border: none;
    width: 100px;
    height: 100px;
  }
  .n0 {
    grid-area: n0;
  }
  .n1 {
    grid-area: n1;
  }
  .n2 {
    grid-area: n2;
  }
  .n3 {
    grid-area: n3;
  }
  .n4 {
    grid-area: n4;
  }
  .n5 {
    grid-area: n5;
  }
  .n6 {
    grid-area: n6;
  }
  .n7 {
    grid-area: n7;
  }
  .n8 {
    grid-area: n8;
  }
  .n9 {
    grid-area: n9;
  }
  .del {
    background:url(./assets/del.png) no-repeat;
    background-size: contain;
    width: 90px;
    height: 80px;
    font-size: 0;
    grid-area: del;
  }
  .ok {
    background:url(./assets/ok.png) no-repeat;
    background-size: contain;
    width: 110px;
    height: 110px;
    font-size: 0;
    grid-area: ok;
  }
  .paid {
    background:url(./assets/confirm.png) no-repeat;
    background-size: contain;
    background-position: center;
    margin-top: 10px;
  }

</style>
