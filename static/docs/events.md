# Events Drive Behaviour
In `staci`, you can create an event using `staci.event("eventName", callback)`. Events can then be registered to elements using the `st-event-type="eventName"` syntax.

## Creating an Event
Here, we establish a `signal` and create an `event` to increment the `signal`:
```html
<script>
  staci.set("count", 0);
  staci.event("incrementCount", () => {
    let [count, setCount] = staci.get("count");
    setCount(count+1);
  });
</script>
```

## Registering an Event
Once an event has been created, you can register it like so:
```html
<button st-click="incrementCount">Click Me!</button>
```

## Multiple Events
Multiple events can be registered by using a semicolor `;` as a seperator:
```html
<button st-click="incrementCount;anotherEvent">Click Me!</button>
```

## Available Events
Here is an array of available `st-` event associations:
```js
this.eventPairs = [
    ["st-click", "click"],
    ["st-dblclick", "dblclick"],
    ["st-mouseup", "mouseup"],
    ["st-mousedown", "mousedown"],
    ["st-mousemove", "mousemove"],
    ["st-mouseenter", "mouseenter"],
    ["st-mouseleave", "mouseleave"],
    ["st-keydown", "keydown"],
    ["st-keypress", "keypress"],
    ["st-keyup", "keyup"],
    ["st-focus", "focus"],
    ["st-blur", "blur"],
    ["st-submit", "submit"],
    ["st-change", "change"],
    ["st-input", "input"],
    ["st-focusin", "focusin"],
    ["st-focusout", "focusout"],
    ["st-select", "select"],
    ["st-resize", "resize"],
    ["st-scroll", "scroll"],
    ["st-wheel", "wheel"],
    ["st-touchstart", "touchstart"],
    ["st-touchend", "touchend"],
    ["st-touchmove", "touchmove"],
    ["st-touchcancel", "touchcancel"],
    ["st-pointerdown", "pointerdown"],
    ["st-pointerup", "pointerup"],
    ["st-pointermove", "pointermove"],
    ["st-pointerenter", "pointerenter"],
    ["st-pointerleave", "pointerleave"],
    ["st-pointercancel", "pointercancel"],
    ["st-contextmenu", "contextmenu"],
    ["st-wheel", "wheel"],
    ["st-drag", "drag"],
    ["st-dragstart", "dragstart"],
    ["st-dragend", "dragend"],
    ["st-dragover", "dragover"],
    ["st-dragenter", "dragenter"],
    ["st-dragleave", "dragleave"],
    ["st-drop", "drop"],
    ["st-input", "input"],
    ["st-keydown", "keydown"],
    ["st-keyup", "keyup"],
];
```
