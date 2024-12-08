# What are Signals?
Signals are points of data located somewhere within the DOM. When a signal's value changes, the value in the DOM will reflect the change automatically.

Put plainly, signals allow us to bypass using `document.querySelector` and related methods to handle DOM updates.

Signals allow to go from a fine-grained updates like this:
```html
<script>
    let button = document.getElementById('btn');
    button.addEventListener('click', () => {
        let text = document.getElementById('text');
        if (text.classList.contains('text-blue-500')) {
            text.classList.remove('text-blue-500');
            text.classList.add('text-red-500');
        } else {
            text.classList.remove('text-red-500');
            text.classList.add('text-blue-500');
        }
    });
</script>

<button id='btn'>Click!</button>
<p id='text' class="text-blue-500">Change my color by clicking</p>
```

To a more reactive model like this:
```html
    <script>
        staci.set('colorClass', 'text-blue-500')
        staci.event('changeColor', () => {
            let [color, setColor] = staci.get('colorClass')
            if (color == 'text-blue-500') {
                setColor('text-red-500')
            } else {
                setColor('text-blue-500')
            }
        })
    </script>

    <button st-click='changeColor'>Click!</button>
    <p class="{| colorClass |}">Change my color by clicking</p>
```

## Creating a Signal
To create a signal:
```html
<script>
    staci.set("count", 0)
</script>
```

## Mounting a Signal Within the DOM
To use the signal's value within the DOM, use the `<st-signal>` custom element.
```html
<script>
    staci.set("count", 0)
</script>
<p>
    <st-signal>{| count |}</st-signal>
</p>
```

`<st-signal>` is used to ensure the elements value does not flicker placeholders `{{}}` prior to `staci` loading. Also, `<st-signal>` isolates the value from the rest of the text in the elements body so it is more confidently updated.

## Mounting a Signal Within an Attribute
To use a signal's value within an element's attribut, do this:
```html
<script>
    staci.set("display", "hidden")
</script>
<p class="{| display |}">I am hidden!</p>
```