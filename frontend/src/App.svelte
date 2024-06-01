<script>
  import {DigitNumber, Clear, Backspace, ClearMemory} from '../wailsjs/go/main/App.js'

  let resultText = "0"
  let memory = "100+"
  $: console.log(resultText)

  function digitNumber(digit) {
    DigitNumber(digit, resultText).then(result => {
      resultText = result
    })
  }

  function c_clear() {
    Clear().then(result => {
      resultText = result
    })
  }

  function ce_clear() {
    c_clear()
    ClearMemory().then(result => {
      memory = result
    })
  }

  function backspace() {
    Backspace(resultText).then(result => {
      resultText = result
    })
  }

  
</script>

<main>
  <div class="calculator">
    <input class="memory" type="text" disabled bind:value={memory} />
    <input class="result" type="text" disabled bind:value={resultText} />
    <button class="button">%</button>
    <button class="button" on:click={() => ce_clear()} >CE</button>
    <button class="button" on:click={() => c_clear()}>C</button>
    <button class="button" on:click={() => backspace()}>⌫</button>
    <button class="button">&#8543;x</button>
    <button class="button">x&#178;</button>
    <button class="button">&#8730;x</button>
    <button class="button operator">%</button>
    <button class="button" on:click={() => digitNumber(7)}>7</button>
    <button class="button" on:click={() => digitNumber(8)}>8</button>
    <button class="button" on:click={() => digitNumber(9)}>9</button>
    <button class="button operator">*</button>
    <button class="button" on:click={() => digitNumber(4)}>4</button>
    <button class="button" on:click={() => digitNumber(5)}>5</button>
    <button class="button" on:click={() => digitNumber(6)}>6</button>
    <button class="button operator">-</button>
    <button class="button" on:click={() => digitNumber(1)}>1</button>
    <button class="button" on:click={() => digitNumber(2)}>2</button>
    <button class="button" on:click={() => digitNumber(3)}>3</button>
    <button class="button operator">+</button>
    <button class="button operator">+/-</button>
    <button class="button" on:click={() => digitNumber(0)}>0</button>
    <button class="button">.</button>
    <button class="button operator">=</button>
  </div>
</main>

<style>
  .result {
      text-align: right;
      padding: 0 10px;
      font-size: 1.5em;
  }
  .calculator { 
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      grid-gap: 2px;
  }
  .button {
      width: 100%;
      height: 50px;
      border-radius: 10%;
      background-color: rgb(70, 70, 70);
      text-align: center;
      font-size: 1.2em;
      color: white;
      cursor: pointer;
      &:hover {
          background-color: rgb(60, 60, 60);
      }
  }
  .operator {
      background-color: #ff9800;
      color: white;
  }
  .memory {
      color: grey;
      background-color: rgb(32, 32, 32);
      border: none;
      grid-column: span 4;
      font-size: 14px;
      text-align: right;
      margin-top:  20px;
  }
  .result {
      grid-column: span 4;
      color: white;
      background-color: rgb(32, 32, 32);
      border: none;
      margin-bottom: 20px;
      font-size: 30px;
  }

 


</style>
