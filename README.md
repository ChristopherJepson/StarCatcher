# StarCatcher: 2D Physics & Particle Study
### Defold Engine | Lua | Cross-Platform Build Pipeline

<div align="center">
  <img src="https://i.ibb.co/DPr6R6r9/Star-Catcher-v0-0-6-screenshot01.png" 
  width="100%" height="auto" alt="StarCatcher Gameplay Screenshot">
</div>

![Defold](https://img.shields.io/badge/Engine-Defold-blue?style=flat&logo=defold)
![Language](https://img.shields.io/badge/Language-Lua%205.1-000080?style=flat&logo=lua)
![Status](https://img.shields.io/badge/Status-Tech%20Demo-yellow)

> **Engineering Context:** This project is a deliberate engine evaluation — 
> assessing Defold's architecture, build pipeline, and cross-platform 
> compilation capabilities against Unity and Unreal Engine 5. A Pipeline 
> Engineer needs to understand multiple build toolchains. This is that work.

---

## 🔧 Technical Analysis: Defold vs. Unity

The primary goal was to translate standard gameplay loops into Defold's 
component architecture within a compressed timeline, evaluating the 
architectural tradeoffs relevant to build pipeline integration.

### 1. Message-Passing Architecture

Unlike Unity's direct method calls (e.g., `GetComponent<T>().DoSomething()`),
this project utilizes Defold's **Message Passing** system for entity 
communication — a fundamentally different decoupling model with direct 
implications for how game logic is structured and tested in a CI pipeline.

- **Hash-Based Logic:** Implemented `msg.post()` systems to handle collisions 
  and state changes, ensuring strict decoupling between game objects.
- **Addressable Components:** Logic relies on URL-based addressing schemes to 
  locate game object instances at runtime.

### 2. Cross-Platform Build Pipeline

- **Multi-Target Compilation:** Configured the Defold render pipeline to 
  target **Windows (x64)** and **macOS (ARM64/x64)** from a single Lua 
  codebase — verifying cross-platform build stability and identifying 
  platform-specific packaging considerations.
- **Build Artifact Management:** Produced versioned release artifacts per 
  platform, validating the build output before distribution.

### 3. Tech Art & Rendering

- **Particle Integration:** Implemented Defold's native particle fx components 
  for visual feedback, analyzing 2D sprite performance overhead under 
  a focused timeline.
- **Atlas Management:** Utilized texture atlases to optimize draw calls, 
  ensuring efficient rendering on lower-end devices.

---

## 🤖 Methodology: AI-Accelerated Development

**This project was executed using an AI-Augmented Workflow.**

- **Rapid Adaptation:** Leveraged Generative AI (Gemini) to reduce the 
  learning curve of the Defold API — implementing features in days rather 
  than weeks.
- **Concept Mapping:** AI tools translated known engineering concepts 
  (State Machines, Vector Math) into Lua syntax, acting as a real-time 
  documentation bridge. Architectural logic remained the human responsibility; 
  AI handled syntactical translation.

---

## 📦 Releases & Downloads

*Current Version: v0.0.6*

| Version | Platform | Status |
| :--- | :--- | :--- |
| **[v0.0.6](https://github.com/ChristopherJepson/StarCatcher/releases/download/v0.0.6/StarCatcher.0.0.6.zip)** | **Windows (x64)** | Stable. High Score tracking added. |
| [v0.0.6](https://github.com/ChristopherJepson/StarCatcher/releases/download/v0.0.6/StarCatcher_MacOS.0.0.6.zip) | macOS (Universal) | Untested. |
| [v0.0.2](https://github.com/ChristopherJepson/StarCatcher/releases/download/v0.0.2/StarCatcher.0.0.2.zip) | Windows (x64) | Prototype. Basic movement only. |

---

## 👤 Author

**Christopher Jepson**  
*Build & Tools Engineer*  
[LinkedIn](https://www.linkedin.com/in/christopher-jepson-310a84308) | 
[Email](mailto:christopher.j.jepson@gmail.com)
