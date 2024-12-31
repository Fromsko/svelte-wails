<script>
  import { createEventDispatcher } from 'svelte'

  const dispatch = createEventDispatcher()

  function sayHello() {
    dispatch('message', {
      text: 'Hello!',
    })
  }
</script>

<button on:click={sayHello}> Outer-Comp Say </button>
