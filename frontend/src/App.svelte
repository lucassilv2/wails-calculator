<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet} from '../wailsjs/go/main/App.js'

  let resultText = "0"
  let name

  function greet() {
    Greet(name).then(result => resultText = result)
  }
</script>

<main>
  <div class="calculator">
    <input class="result" type="text" bind:value={resultText} />
    <button class="button">%</button>
    <button class="button">CE</button>
    <button class="button">C</button>
    <button class="button">⌫</button>
    <button class="button">&#8543;x</button>
    <button class="button">x&#178;</button>
    <button class="button">&#8730;x</button>
    <button class="button operator">%</button>
    <button class="button">7</button>
    <button class="button">8</button>
    <button class="button">9</button>
    <button class="button operator">*</button>
    <button class="button">4</button>
    <button class="button">5</button>
    <button class="button">6</button>
    <button class="button operator">-</button>
    <button class="button">1</button>
    <button class="button">2</button>
    <button class="button">3</button>
    <button class="button operator">+</button>
    <button class="button operator">+/-</button>
    <button class="button">0</button>
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
  .result {
      grid-column: span 4;
      color: white;
      background-color: rgb(32, 32, 32);
      border: none;
      height: 70px;
      margin-bottom: 5px;
  }

</style>
