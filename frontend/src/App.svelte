<script>
  import {AGenerateKeyPair, AEncryptMessage, ADecryptMessage} from '../wailsjs/go/main/App.js'

  let msgText = "";
  let publKey = "";
  let privKey = "";
  let defaultMsg = "";
  let cryptMsg = "";

  const generateKeys = () => {
    AGenerateKeyPair().then(publAndPriv => {
      publKey = publAndPriv[0];
      privKey = publAndPriv[1];
    });
  }

  const encryptMessage = () => {
    if (publKey === "") {
      msgText = "Публічний ключ не повинен бути порожнім!"
      return;
    }
    if (defaultMsg === "") {
      msgText = "Введіть щось"
      return;
    }
    AEncryptMessage(defaultMsg, publKey).then(res => {
      if (res[1] !== "0") {
        msgText = res[1];
        return;
      }
      cryptMsg = res[0];
      defaultMsg = "";
    });
  }

  const decryptMessage = () => {
    if (privKey === "") {
      msgText = "Приватний ключ не повинен бути порожнім!"
      return;
    }
    if (cryptMsg === "") {
      msgText = "Введіть щось"
      return;
    }
    ADecryptMessage(cryptMsg, privKey).then(res => {
      if (res[1] !== "0") {
        msgText = res[1];
        return;
      }
      defaultMsg = res[0];
      cryptMsg = "";
    });
  }
</script>

<main>
  {#if msgText !== ""}
  <div class="overlay">
    <div class="alert">
      <p>Сповіщення</p>
      <div class="content">{msgText}</div>
      <button on:click={() => {msgText="";}}>OK</button>
    </div>
  </div>
  {/if}
  <h1>Welcome to AdvRsa! 🔑</h1>

  <div class="two-columns">
    <div class="column">
      <p>Публічний ключ 2048bit (base64):</p>
      <textarea class="publKey" bind:value={publKey}></textarea>
    </div>
    <div class="column">
      <p>Приватний ключ 2048bit (base64):</p>
      <textarea class="privKey" bind:value={privKey}></textarea>
    </div>
  </div>
  <button class="generate" on:click={generateKeys}>Генерувати ключі</button>

  <div class="two-columns">
    <div class="column">
      <textarea placeholder="Звичайне повідомлення" bind:value={defaultMsg}></textarea>
      <button on:click={encryptMessage}>Зашифрувати</button>
    </div>
    <div class="column">
      <textarea placeholder="Зашифроване повідомлення" bind:value={cryptMsg}></textarea>
      <button on:click={decryptMessage}>Розшифрувати</button>
    </div>
  </div>
</main>

<style>
  p {
    margin: 0;
  }

  .overlay {
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(255, 255, 255, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .alert {
    display: flex;
    flex-direction: column;
    width: 450px;
    background-color: white;
    border: 1px solid black;
    border-radius: 5px;
    padding: 5px;
    box-sizing: border-box;
    padding-bottom: 20px;
    filter: drop-shadow(0px 0px 28px 0px rgba(0,0,0,0.56));
    -webkit-filter: drop-shadow(0px 0px 28px 0px rgb(0,0,0));
    animation: startAlert ease-in 0.2s 0s;
  }

  @keyframes startAlert {
    from {
      transform: scaleY(0);
    }
    to {
      transform: scaleY(1);
    }
  }

  .alert > p {
    font-weight: 200;
    color: #444;
    border-bottom: 1px solid black;
    padding-bottom: 5px;
    margin-bottom: 10px;
  }

  .alert .content {
    height: 100%;
  }

  main {
    background: linear-gradient(0deg, rgba(255,255,255,1) 0%, rgba(216,216,255,1) 32%, rgba(94,228,255,1) 100%);
    box-sizing: border-box;
    position: absolute;
    padding: 15px;
    width: 100%;
    height: 100%;
  }

  h1 {
    background-color: rgb(255, 255, 255, 0.5);
    font-weight: 400;
    text-align: center;
    margin: 0;
    border-radius: 5px;
  }

  .two-columns {
    padding-top: 30px;
    width: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
  }

  .column > * {
    box-sizing: border-box;
    width: 100%;
  }

  .column > p {
    text-align: center;
    padding-bottom: 10px;
  }

  .column > textarea {
    padding: 5px;
    border-radius: 5px;
    height: 70px;
  }

  button {
    margin-top: 10px;
    width: 100%;
    border-radius: 5px;
    background-color: white;
    border: 1px solid black;
  }

  button:hover {
    background-color: #eee;
  }
</style>
