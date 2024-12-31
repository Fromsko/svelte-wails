<script>
  function handleClick() {
    console.log('clicked')
  }
</script>

<button on:click|once={handleClick}> Click me </button>

<!-- 
所有修饰符列表：

preventDefault ：调用event.preventDefault() ，在运行处理程序之前调用。比如，对客户端表单处理有用。
stopPropagation ：调用 event.stopPropagation(), 防止事件影响到下一个元素。
passive ： 优化了对 touch/wheel 事件的滚动表现(Svelte 会在合适的地方自动添加滚动条)。
capture — 在 capture 阶段而不是bubbling 阶段触发事件处理程序 ()
once ：运行一次事件处理程序后将其删除。
self — 仅当 event.target 是其本身时才执行。

你可以将修饰符组合在一起使用，例如：on:click|once|capture={...}。
-->
