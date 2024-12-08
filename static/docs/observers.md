# I'm Watching You
`staci` comes shipped with a few attribute `observers` to make life easier for common tasks. You can place signals within these attributes to modify the behaviour of elements.

## st-cloak
`st-cloak` allows us to determine an element's [visibility](https://developer.mozilla.org/en-US/docs/Web/CSS/visibility). Note: this is different than an element's display.

Use `st-cloak` when you want to hide an element from view, but *wish to perserve* it's place in the overall layut of the page.

<div class='flex flex-col gap-4 pb-4'>
    <script>
        staci.set('isCloaked', true);
        staci.event('cloakToggle', () => {
            let [isCloaked, setIsCloaked] = staci.get('isCloaked');
            setIsCloaked(!isCloaked);
        })
    </script>
    <button st-click='cloakToggle' class='border border-gray-800 text-gray-400 w-fit px-2 py-1 rounded text-sm'>Click to Toggle</button>
    <p st-cloak="{| isCloaked |}" class='p-4 rounded text-sm border border-gray-800 text-gray-400'>My placement was perserved!</p>
</div>

The code for the above example:
```html
<div>
    <script>
        staci.set('isCloaked', true);
        staci.event('cloakToggle', () => {
            let [isCloaked, setIsCloaked] = staci.get('isCloaked');
            setIsCloaked(!isCloaked);
        })
    </script>
    <button st-click='cloakToggle'>Click to Toggle</button>
    <p st-cloak="{| isCloaked |}">My placement was perserved!</p>
</div>
```

## st-hide
`st-hide` allows us to determine an elements [display](https://developer.mozilla.org/en-US/docs/Web/CSS/display). Note: this is different than an elements visibility.

Use `st-hide` when you want an element, along with its placement in the DOM (its padding, margin, ect), to be removed from view.

<div class='flex flex-col gap-4 pb-4'>
    <script>
        staci.set('isHidden', true);
        staci.event('hideToggle', () => {
            let [isHidden, setIsHidden] = staci.get('isHidden');
            setIsHidden(!isHidden);
        })
    </script>
    <button st-click='hideToggle' class='border border-gray-800 text-gray-400 w-fit px-2 py-1 rounded text-sm'>Click to Toggle</button>
    <p st-hide="{| isHidden |}" class='p-4 rounded text-sm border border-gray-800 text-gray-400'>My placement was not perserved!</p>
</div>

The code for the above example:
```html
<div>
    <script>
        staci.set('isHidden', true);
        staci.event('hideToggle', () => {
            let [isHidden, setIsHidden] = staci.get('isHidden');
            setIsHidden(!isHidden);
        })
    </script>
    <button st-click='hideToggle'>Click to Toggle</button>
    <p st-hide="{| isHidden |}">My placement was not perserved!</p>
</div>
```
