# Staci
drop-in, reactive signals 🤌

## Counter Example
<button class='bg-black border border-gray-800 w-fit rounded px-2 py-1 mb-4 text-sm text-gray-400' st-click='increment-count'>
    <script>
        staci.set("count", 0);
        staci.event("increment-count", () => {
            let [count, setCount] = staci.get("count");
            setCount(count+1);
        });
    </script>
    <span class='pr-4'>Increment</span>
    <st-signal st-force='true'>{| count |}</st-signal>
</button>

```html
<script>
    staci.set("count", 0);
    staci.event("incrementCount", () => {
        let [count, setCount] = staci.get("count");
        setCount(count+1);
    });
</script>

<button st-click="incrementCount">
    <p>Increment</p>
    <st-signal>{| count |}</st-signal>
</button>
```

## Use-Case
`staci` is all about *compliementing* existing HTML. Other solutions dictate the way you write your HTML. When you want to generate HTML, you must do so **with the framework**.

`staci` is different. She is not about replacing you're existing wife__ I mean system. Rather, `staci` intends to join-in where needed and can be adopted incrementally overtime.

## Philosophy
Imagine if signals were available natively in the DOM. That is the vibe `staci` attempts to create.

Using signals should feel *natural*.

![Natural](/static/img/buddy.jpg)

