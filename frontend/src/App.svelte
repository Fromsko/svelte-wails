<script lang="ts">
  import AuthState from './components/auth-if/AuthState.svelte'
  import GenRandom from './components/await/GenRandom.svelte'
  import BindChekbox from './components/bind/BindChekbox.svelte'
  import BindContent from './components/bind/BindContent.svelte'
  import BindInput from './components/bind/BindInput.svelte'
  import BindKeypad from './components/bind/BindKeypad.svelte'
  import BindMulitipe from './components/bind/BindMulitipe.svelte'
  import BindMultArray from './components/bind/BindMultArray.svelte'
  import BindRadio from './components/bind/BindRadio.svelte'
  import BindThis from './components/bind/BindThis.svelte'
  import ParentSun from './components/comp-event/ParentSun.svelte'
  import DomEvent from './components/dom-event/DomEvent.svelte'
  import Cats from './components/each/Cats.svelte'
  import EachComp from './components/each/EachComp.svelte'
  import Info from './components/exports/Info.svelte'
  import DomCompDispatcher from './components/outer-comp-event/dom-dispatcher/DomCompDispatcher.svelte'
  import OutComp from './components/outer-comp-event/inner-dispatcher/OutComp.svelte'

  // import {Greet} from '../wailsjs/go/main/App.js'
  const dragLayout = {
    '--wails-draggable': 'drag',
  }

  const pkg = {
    name: 'svelte',
    version: 3,
    speed: 'blazing',
    website: 'https://svelte.dev',
  }
</script>

<main class="p-6 text-center h-screen rounded-lg">
  <Info {...pkg} />
  <br />
  <AuthState />
  <br />
  <Cats />
  <br />
  <EachComp />
  <br />
  <GenRandom />
  <br />
  <DomEvent />
  <br />
  <ParentSun />
  <br />
  <OutComp />
  <br />
  <DomCompDispatcher />
  <br />
  <BindInput />
  <br />
  <BindRadio />
  <br />
  <BindChekbox />
  <br />
  <BindMultArray />
  <br />
  <BindMulitipe />
  <br />
  <BindContent />
  <br />
  <BindThis />
  <br />
  <BindKeypad />
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme('colors.blue.100');
  }
</style>
