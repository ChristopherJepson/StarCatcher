# StarCatcher: 2D Physics & Particle Study
### Defold Engine | Lua | Cross-Platform Build

<div align="center">
  <img src="https://i.ibb.co/DPr6R6r9/Star-Catcher-v0-0-6-screenshot01.png" width="100%" height="auto" alt="StarCatcher Gameplay Screenshot">
</div>

![Defold](https://img.shields.io/badge/Engine-Defold-blue?style=flat&logo=defold)
![Language](https://img.shields.io/badge/Language-Lua%205.1-000080?style=flat&logo=lua)
![Status](https://img.shields.io/badge/Status-Tech%20Demo-yellow)

> **Engineering Context:** This project serves as a technical evaluation of the **Defold Game Engine**. It demonstrates the ability to rapidly adapt to a "Message-Passing" architecture, contrasting with standard Object-Oriented patterns found in Unity.

---

## 🔧 Technical Analysis: Defold vs. Unity
The primary goal of this repository was to translate standard gameplay loops into Defold's component architecture within a compressed timeline.

### 1. Message-Passing Architecture
Unlike Unity's direct method calls (e.g., `GetComponent<T>().DoSomething()`), this project utilizes Defold's **Message Passing** system to handle entity communication.
- **Hash-Based Logic:** Implemented `msg.post()` systems to handle collisions and state changes, ensuring strict decoupling between game objects.
- **Addressable Components:** Logic relies on URL-based addressing schemes to locate game object instances at runtime.

### 2. Tech Art & Particles
- **Particle Integration:** Implemented Defold's native particle fx components for visual feedback (collection of stars), analyzing the performance overhead of 2D sprites in a focused timeline.
- **Atlas Management:** Utilized texture atlases to optimize draw calls, ensuring efficient rendering on lower-end devices.

### 3. Cross-Platform Compilation
- **Build Pipeline:** Successfully configured the render pipeline to target both **Windows (x64)** and **macOS (ARM64/x64)**, verifying the cross-platform stability of the Lua codebase.

---

## 🤖 Methodology: AI-Accelerated Development
**This project was executed using an AI-Augmented Workflow to maximize efficiency.**
- **Rapid Adaptation:** Leveraged Generative AI (Gemini) to drastically reduce the learning curve of the Defold API. This allowed for the implementation of complex features in days rather than weeks.
- **Concept Mapping:** AI tools were used to translate known Engineering concepts (State Machines, Vector Math) into specific **Lua syntax**, acting as a real-time documentation bridge. The focus remained on **Architectural Logic**, while AI handled the syntactical heavy lifting.

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
*Technical Artist & Software Engineer*
[LinkedIn](https://www.linkedin.com/in/christopher-jepson-310a84308) | [Email](mailto:christopher.j.jepson@gmail.com)
