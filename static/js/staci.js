// contains attribute observers which make making changes to the DOM easier
class Watcher {
  static observerStCloak(element) {
    let callback = (element) => {
      let stHideAttr = element.getAttribute("st-cloak");
      if (stHideAttr == "true") {
        element.style.visibility = "hidden";
      }
      if (stHideAttr == "false") {
        element.style.visibility = "visible";
      }
    };
    callback(element);
    const observer = new MutationObserver(() => {
      callback(element);
    });
    observer.observe(element, {
      attributes: true,
      attributeFilter: ["st-cloak"],
    });
  }
  static observerStHide(element) {
    const callback = (element) => {
      let stHideAttr = element.getAttribute("st-hide");
      if (stHideAttr === "true") {
        // Stash the current display style if not already stashed
        if (!element.hasAttribute("st-hide-prev")) {
          element.setAttribute("st-hide-prev", element.style.display || "");
        }
        element.style.display = "none";
      } else if (stHideAttr === "false") {
        // Restore the stashed display style
        let previousDisplay = element.getAttribute("st-hide-prev");
        element.style.display = previousDisplay || "flex"; // Default to "block" if no stashed value
      }
    };

    // Call the callback initially to handle the current state
    callback(element);

    // Create a MutationObserver to watch for changes in the "st-hide" attribute
    const observer = new MutationObserver(() => {
      callback(element);
    });

    // Start observing the element for changes to the "st-hide" attribute
    observer.observe(element, {
      attributes: true,
      attributeFilter: ["st-hide"],
    });
  }


}

// contains wrapper functions like throttle and debounce to wrap events
class Everest {
  static throttle(func, delay) {
    let lastTime = 0;
    return function (...args) {
      const now = Date.now();
      if (now - lastTime >= delay) {
        lastTime = now;
        func.apply(this, args);
      }
    };
  }

  static debounce(func, delay) {
    let timeout;
    return function (...args) {
      clearTimeout(timeout);
      timeout = setTimeout(() => func.apply(this, args), delay);
    };
  }
}

// contains functions used to access and manipulate the dom
class Dom {
  static getDirectTextContent(element) {
    let directText = "";
    element.childNodes.forEach((node) => {
      if (node.nodeType === Node.TEXT_NODE) {
        directText += node.nodeValue;
      }
    });
    return directText.trim();
  }

  static replaceTextContent(element, placeholder, newValue) {
    const currentText = element.textContent;
    const updatedText = currentText.replace(placeholder, newValue);
    element.textContent = updatedText;
  }

  static ClimbDomUntil(element, callback) {
    if (!element || typeof callback !== "function") {
      throw new Error(
        "Invalid arguments: Provide a valid element and a callback function.",
      );
    }
    let currentElement = element;
    while (currentElement) {
      if (callback(currentElement)) {
        return true;
      }
      if (currentElement === document.body) {
        break;
      }
      currentElement = currentElement.parentElement;
    }
    return false;
  }
}

// contains functions intended to be used to iterate over different types of data
class Iter {
  static map(items, fn) {
    for (let i = 0; i < items.length; i++) {
      let shouldContinue = fn(items[i]);
      if (shouldContinue == null || shouldContinue == undefined) {
        console.error(
          "forgot to return true from Iter.map(items, fn), return false to break",
        );
      }
      if (!shouldContinue) {
        break;
      }
    }
  }

  static mapObj(obj, fn) {
    for (let [key, value] of Object.entries(obj)) {
      let shouldContinue = fn(key, value);
      if (shouldContinue == null || shouldContinue == undefined) {
        console.error(
          "forgot to return true from Iter.mapObj(items, fn), return false to break",
        );
      }
      if (!shouldContinue) {
        break;
      }
    }
  }
}

// contains string parsing functions
class Purse {
  static scanBetweenSubStrs(s, start, end) {
    const out = [];
    let inSearch = false;
    let searchStr = "";
    let i = 0;

    while (i < s.length) {
      if (
        !inSearch &&
        i + start.length <= s.length &&
        s.substring(i, i + start.length) === start &&
        (i === 0 || s[i - 1] !== "\\") // Check for preceding escape character
      ) {
        inSearch = true;
        searchStr = start;
        i += start.length;
        continue;
      }

      if (inSearch) {
        if (
          i + end.length <= s.length &&
          s.substring(i, i + end.length) === end
        ) {
          searchStr += end;
          out.push(searchStr);
          searchStr = "";
          inSearch = false;
          i += end.length;
          continue;
        }
        searchStr += s[i];
      }

      i++;
    }
    return out;
  }

  static removeFromArray(array, ...valuesToRemove) {
    return array.filter((item) => !valuesToRemove.includes(item));
  }

  static strIsInt(str) {
    if (typeof str !== "string") return false; // Ensure the input is a string
    const num = Number(str);
    return Number.isInteger(num) && str.trim() !== "";
  }

  static ifElse(value, potential1, potential2) {
    if (value == potential1) {
      return potential2;
    }
    if (value == potential2) {
      return potential1;
    }
    console.error("used Purse.ifElse, but function returned a null value");
    return null;
  }
}

// a class to handle signals within our application
class Signal {
  constructor(initialValue) {
    this.value = initialValue;
    this.listeners = new Set();
    this.history = [initialValue];
  }

  val() {
    return this.value;
  }

  set(newValue) {
    let oldVal = this.val();
    if (this.value !== newValue) {
      this.value = newValue;
      this.notify(oldVal);
    }
  }

  subscribe(callback) {
    this.listeners.add(callback);
    return () => this.listeners.delete(callback); // Return an unsubscribe function
  }

  notify(oldVal) {
    this.listeners.forEach((callback) => callback(oldVal, this.value));
  }
}

// provides a st-scollbar class to custom style a scrollbar
class CustomScroll extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    this.innerHTML = `
            <style>
                .custom-scroll::-webkit-scrollbar {
                    width: 8px;
                    height: 8px;
                }
                .custom-scroll::-webkit-scrollbar-thumb {
                    background-color: #4B5563; /* Gray-600 */
                    border-radius: 4px;
                }
                .custom-scroll::-webkit-scrollbar-track {
                    background-color: #1F2937; /* Gray-800 */
                }
            </style>
        `;
  }
}

class TitleLinks extends HTMLElement {
    constructor() {
      super();
    }

    connectedCallback() {
      const targetSelector = this.getAttribute('target');
      const offset = parseInt(this.getAttribute('offset'), 10) || 0;
      
      const targetElement = document.querySelector(targetSelector);
      if (!targetElement) {
        console.error(`Target element "${targetSelector}" not found.`);
        return;
      }
      
      const headings = targetElement.querySelectorAll('h1, h2, h3, h4, h5, h6');

      headings.forEach(heading => {
        if (heading.id) {
          const linkItem = document.createElement('div');
          const link = document.createElement('a');
          link.classList.add('title-link');
          link.href = `#${heading.id}`;
          link.textContent = heading.textContent;
          linkItem.appendChild(link);
          this.appendChild(linkItem);
        }
      });

      // Add styles
      const style = document.createElement('style');
      this.appendChild(style);

      // Scroll adjustment with offset when clicking links
      this.addEventListener('click', (e) => {
        if (e.target.tagName === 'A') {
          e.preventDefault();
          const targetId = e.target.getAttribute('href').substring(1);
          const targetElement = document.getElementById(targetId);
          if (targetElement) {
            const position = targetElement.getBoundingClientRect().top + window.pageYOffset + offset;
            window.scrollTo({
              top: position,
              behavior: 'smooth'
            });
          }
        }
      });
    }
  }

// staci uses this custom element to determine the placement of signals within the DOM
class StSignal extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.setAttribute("style", "visibility:hidden;");
  }
}

// registering all of our custom elements
customElements.define("custom-scroll", CustomScroll);
customElements.define("st-signal", StSignal);
customElements.define("title-links", TitleLinks);

// a class to handle dynamic web interactions
class Staci {
  constructor() {
    this.version = "0.0.1";
    this.repo = "https://github.com/phillip-england/staci";
    this.events = {};
    this.signals = {};
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
    this.initOnLoad();
  }

  initOnLoad() {
    document.addEventListener("DOMContentLoaded", () => {
      this.initAllEventTypes();
      this.initSignalTextPlaceholders();
      this.initSignalAttrPlaceholders();
      this.initAttrObservers();
    });
  }

  initAllEventTypes() {
    Iter.map(this.eventPairs, (pair) => {
      let stEvent = pair[0];
      let htmlEvent = pair[1];
      this.initEventsOfType(stEvent, htmlEvent);
      return true;
    });
  }

  initEventsOfType(eventAttr, mapsTo) {
    Iter.mapObj(this.events, (key, event) => {
      let subElms = document.querySelectorAll(`*[${eventAttr}]`);
      Iter.map(subElms, (elm) => {
        let sublist = elm.getAttribute(`${eventAttr}`);
        let parts = sublist.split(";");
        Iter.map(parts, (part) => {
          let event = staci.getEvent(part);
          let elmThrottle = elm.getAttribute("st-throttle");
          if (elmThrottle) {
            let throttleIsInt = Purse.strIsInt(elmThrottle);
            if (!throttleIsInt) {
              console.error(
                `element ${element} contains st-throttle with a string that cannot be converted to an int ${elmThrottle}`,
              );
            }
            let throttle = parseInt(elmThrottle);
            let currentEvents = elm.getAttribute("st-events");
            if (!currentEvents) {
              elm.addEventListener(mapsTo, Everest.throttle(event, throttle));
              elm.setAttribute("st-events", eventAttr);
            } else {
              if (!currentEvents.includes(eventAttr)) {
                elm.addEventListener(mapsTo, Everest.throttle(event, throttle));
                elm.setAttribute("st-events", currentEvents + ";" + eventAttr);
              }
            }
            return true;
          }
          let elmDebounce = elm.getAttribute("st-debounce");
          if (elmDebounce) {
            let debounceIsInt = Purse.strIsInt(elmDebounce);
            if (!debounceIsInt) {
              console.error(
                `element ${element} contains st-debounce with a string that cannot be converted to an int ${elmDebounce}`,
              );
            }
            let debounce = parseInt(elmDebounce);
            let currentEvents = elm.getAttribute("st-events");
            if (!currentEvents) {
              elm.addEventListener(mapsTo, Everest.debounce(event, debounce));
              elm.setAttribute("st-events", eventAttr);
            } else {
              if (!currentEvents.includes(eventAttr)) {
                elm.addEventListener(mapsTo, Everest.debounce(event, debounce));
                elm.setAttribute("st-events", currentEvents + ";" + eventAttr);
              }
            }
            return true;
          }
          let currentEvents = elm.getAttribute("st-events");
          if (!currentEvents) {
            elm.addEventListener(mapsTo, event);

            elm.setAttribute("st-events", eventAttr);
          } else {
            if (!currentEvents.includes(eventAttr)) {
              elm.addEventListener(mapsTo, event);

              elm.setAttribute("st-events", currentEvents + ";" + eventAttr);
            }
          }

          return true;
        });
        return true;
      });
      return true;
    });
  }

  initSignalTextPlaceholders() {
    let allElms = document.querySelectorAll("*");
    let ignoreElms = [];
    Iter.map(allElms, (elm) => {
      let stIgnore = elm.getAttribute("st-ignore");
      if (stIgnore == null) {
        return true;
      }
      if (stIgnore == "true") {
        ignoreElms.push(elm);
      }
      return true;
    });
    Iter.map(allElms, (elm) => {
      let hasStIgnoreParent = Dom.ClimbDomUntil(elm, (parent) => {
        let stIgnoreAttr = parent.getAttribute("st-ignore");
        if (stIgnoreAttr == null) {
          return false;
        }
        if (stIgnoreAttr == "true") {
          return true;
        }
      });
      let stForce = elm.getAttribute("st-force");
      let forceElm = false;
      if (stForce) {
        if (stForce == "true") {
          forceElm = true;
        }
      }
      if (!forceElm && hasStIgnoreParent) {
        return true;
      }
      let text = Dom.getDirectTextContent(elm);
      let placeholders = Purse.scanBetweenSubStrs(text, "{|", "|}");
      Iter.map(placeholders, (placeholder) => {
        let signalKey = placeholder.replace("{|", "").replace("|}", "").trim();
        let signal = staci.getSignalFull(signalKey);
        if (signal == null || signal == undefined) {
          console.error(
            `attempting to set signal values on a {{ }} placeholder, but
                    we encountered a null signal from using the key ${signalKey} and placeholder ${placeholder}`,
          );
        }
        Dom.replaceTextContent(elm, placeholder, signal.val());
        let styleAttr = elm.getAttribute("style");
        if (styleAttr == null) {
          return true;
        }
        if (styleAttr.includes("hidden")) {
          elm.setAttribute("style", "");
        }
        signal.subscribe((oldVal, newVal) => {
          Dom.replaceTextContent(elm, oldVal, newVal);
        });
        return true;
      });
      return true;
    });
  }

  initSignalAttrPlaceholders() {
    let allElms = document.querySelectorAll("*");
    Iter.map(allElms, (elm) => {
      Iter.map(elm.attributes, (attr) => {
        let attrKey = attr.name;
        let attrVal = attr.value;
        let placeholders = Purse.scanBetweenSubStrs(attrVal, "{|", "|}");
        Iter.map(placeholders, (placeholder) => {
          let signalKey = placeholder
            .replace("{|", "")
            .replace("|}", "")
            .trim();
          let hasExclamation = false;
          if (signalKey[0] == "!") {
            signalKey = signalKey.replace("!", "");
            hasExclamation = true;
          }
          let signal = staci.getSignalFull(signalKey);
          if (signal == null || signal == undefined) {
            console.error(
              `attempting to set signal values on a {{ }} placeholder, but
                        we encountered a null signal from using the key ${signalKey} and placeholder ${placeholder}`,
            );
          }
          let callback = null
          if (hasExclamation) {
            attrVal = attrVal.replace(placeholder, !signal.val());
            callback = (oldVal, newVal) => {
                if (typeof(oldVal) == "string" && typeof(newVal) == "string") {
                    let attrParts = attrVal.split(' ')
                    attrParts = attrParts.filter((part) => {
                        if (part == newVal) {
                            return false
                        }
                        return true
                    })
                    attrParts.push(...newVal.split(' '))
                    elm.setAttribute(attrKey, attrParts.join(' '));
                    return
                }
                attrVal = attrVal.replace(!oldVal, !newVal);
                elm.setAttribute(attrKey, attrVal);
            }
          } else {
            attrVal = attrVal.replace(placeholder, signal.val());
            callback = (oldVal, newVal) => {
                if (typeof(oldVal) == "string" && typeof(newVal) == "string") {
                    let attrParts = attrVal.split(' ')
                    attrParts = attrParts.filter((part) => {
                        if (part == oldVal) {
                            return false
                        }
                        return true
                    })
                    attrParts.push(...newVal.split(' '))
                    elm.setAttribute(attrKey, attrParts.join(' '));
                    return
                }
                attrVal = attrVal.replace(oldVal, newVal);
                elm.setAttribute(attrKey, attrVal);
            }
          }
          elm.setAttribute(attrKey, attrVal);
          signal.subscribe(callback);
          return true;
        });
        return true;
      });
      return true;
    });
  }

  initAttrObservers() {
    let allElms = document.querySelectorAll("*");
    Iter.map(allElms, (elm) => {
      Watcher.observerStHide(elm);
      Watcher.observerStCloak(elm);
      return true;
    });
  }

  event(name, fn) {
    this.events[name] = fn;
  }

  set(key, val) {
    this.signals[key] = new Signal(val);
  }

  getEvent(key) {
    let foundEvent = false;
    let event = null;
    Iter.mapObj(this.events, (innerKey, val) => {
      if (key == innerKey) {
        foundEvent = true;
        event = val;
        return false;
      }
      return true;
    });
    if (foundEvent == false) {
      console.error(
        `attempted to access event named ${key} when one does not exist`,
      );
      return null;
    }
    return event;
  }

  get(key) {
    let foundSignal = false;
    let signal = null;
    Iter.mapObj(this.signals, (innerKey, val) => {
      if (key == innerKey) {
        foundSignal = true;
        signal = val;
        return false;
      }
      return true;
    });
    if (foundSignal == false) {
      console.error(
        `attempted to access signal named ${key} when one does not exist`,
      );
      return null;
    }
    return [signal.val(), (newVal) => {
      signal.set(newVal);
    }];
  }

  getSignalFull(key) {
    let foundSignal = false;
    let signal = null;
    Iter.mapObj(this.signals, (innerKey, val) => {
      if (key == innerKey) {
        foundSignal = true;
        signal = val;
        return false;
      }
      return true;
    });
    if (foundSignal == false) {
      console.error(
        `attempted to access signal named ${key} when one does not exist`,
      );
      return null;
    }
    return signal;
  }
}

let staci = new Staci();
console.log(staci);
