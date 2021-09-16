<script>
  import { Router, Link, Route } from "svelte-routing";
  import Todo from "./components/todo/Todo.svelte";

  export let url = "";
</script>

<Router url="{url}">
  <nav>
    <Link to="/">Todo</Link>
  </nav>
  <div>
    <Route path="/"><Todo /></Route>
  </div>
</Router>