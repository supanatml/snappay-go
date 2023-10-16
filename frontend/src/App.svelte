<script>
  import logo from './assets/images/logo-universal.png'
  import termsPNG from './assets/terms.png'
  import {QRpopup} from '../wailsjs/go/main/App.js'
  import {VoucherPopup} from '../wailsjs/go/main/App.js'

  let page = 0

  async function showQR() {
    await QRpopup()
    page = 1
  }
  async function showVoucher() {
    await VoucherPopup()
    page = 2
  }
  function showTerms() {
    page = 3
  }
  function checkTx() {
    page = 0
  }
</script>

<main>
  {#if page === 0}
    <img alt="Wails logo" id="logo" src="{logo}">
    <div class="btn-group">
      <button class="btn start" on:click={showQR}>Start</button>
      <button class="btn voucher" on:click={VoucherPopup}>Voucher</button>
      <button class="btn terms" on:click={showTerms}>Terms</button>
    </div>
  {:else if page === 1}
      <img id="qr" src="qr.png" alt="No QR available"/>
      <button class="btn paid" on:click={checkTx}>Confirm payment</button>
  {:else if page === 2}
        <img id="qr" src="v.png?" alt="No v.png available"/>
  {:else if page === 3}
    <img id="terms" src="{termsPNG}" alt="Terms unavailable" width="100%">
    <button class="btn backbtn" on:click={() => (page = 0)}>Back</button>
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
    margin-bottom: 60px;
  }
  .voucher {
    background:url(./assets/voucherbtn.png) no-repeat;
    background-size: contain;
    background-position: center;
  }
  .terms {
    background:url(./assets/viewterms3.png) no-repeat;
    background-size: contain;
    background-position: center;
  }

  .paid {
    background:url(./assets/paid2.png) no-repeat;
    background-size: contain;
    background-position: center;
  }

  .backbtn {
    width: 60px;
    height: 30px;
    font-size: 1rem;
  }

</style>
