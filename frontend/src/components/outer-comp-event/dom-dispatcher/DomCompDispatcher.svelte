<script>
  import FancyButton from './CustomButton.svelte'

  function handleClick() {
    console.log('clicked')
  }
</script>

<FancyButton on:click={handleClick} />
