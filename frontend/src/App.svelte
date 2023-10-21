<script>
  import {QRpopup} from '../wailsjs/go/main/App.js'
  import {CheckTransaction} from '../wailsjs/go/main/App.js'
  import {VoucherPopup} from '../wailsjs/go/main/App.js'
  import {PrintLog} from '../wailsjs/go/main/App.js'

  let page = 0
  let qrPromise
  let timeoutID
  let ref = "0"
  //todo: pass ref instead of global var?

  function showChoice() {
    page = 1
  }
  async function showQR() {
    page = 2
    ref = await QRpopup();
	  timeoutID = setTimeout(() => page = 0, 600000);
    return ref
  }
  async function showVoucher() {
    page = 3
	  setTimeout(() => page = 0, 60000);
    await VoucherPopup()
  }
  function checkPaid() {
    let paid = CheckTransaction(ref)
    if(paid) {
      clearTimeout(timeoutID)
    }
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
        <img id="qr" src="qr.png?ref={ref}" alt="No QR available"/>
    {/await}
      <button class="btn paid" on:click={checkPaid}>Confirm payment</button>
  {:else if page === 3}
    <div class="voucherpage">
      <button class="btn back" on:click={() => (page = 0)}>Back</button>
      <label class="voucherinput">Input your voucher</label>
      <button class="numpad">0</button>
      <button class="numpad">1</button>
      <button class="numpad del">Del</button>
      <button class="numpad ok">OK</button>
    </div>
  {/if}
</main>

<style>

  #terms {
    display: block;
    width: 100%;
    margin-bottom: 40px;
  }

  #qr {
    display: block;
    width: 100%;
  }

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
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
    position: fixed;
    left: -40px;
    overflow: visible;
  }
  .numpad {
    background:url(./assets/blank.png) no-repeat;
    background-size: contain;
    background-position: center;
    width: 120px;
    height: 120px;
    font-size: 2rem;
    border: none;
  }
  .del {
    background:url(./assets/del.png) no-repeat;
    background-size: contain;
    width: 100px;
    height: 90px;
    font-size: 0;
  }
  .ok {
    background:url(./assets/ok.png) no-repeat;
    background-size: contain;
    width: 120px;
    height: 120px;
    font-size: 0;
  }
  .paid {
    background:url(./assets/confirm.png) no-repeat;
    background-size: contain;
    background-position: center;
    margin-top: 10px;
  }

  .backbtn {
    width: 60px;
    height: 30px;
    font-size: 1rem;
  }

</style>
