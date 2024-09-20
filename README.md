# Communication System for Disaster Monitoring

## Project Objective

This project aims to implement a multi-protocol communication system for disaster monitoring, using ESP32 devices as communication units and a central coordinator to manage communications. The system uses the MQTT protocol to ensure reliable and efficient communication between the units and the coordinator.

---

## Coordinator

### Overview

The coordinator is responsible for managing the communication between the different field units (ESP32) and processing the received data. It receives control instructions, sends "wake-up" signals to the units as needed, and forwards the processed data to the external system.

### MQTT Topics

- **`/Control/Event/#`**: 
  - **Description**: The coordinator subscribes to this topic to monitor control events that may require specific actions, such as instructions to wake up units or change the system state.
  - **Rationale**: This topic allows the coordinator to respond to global events that can affect the operation of the system as a whole.

- **`/Data/Coordinator/#`**:
  - **Description**: A topic used by the coordinator to receive data or specific instructions from the units or other control systems.
  - **Rationale**: Centralizes the receipt of data and instructions for the coordinator, facilitating the processing and dissemination of information to the units.

- **`/Control/WakeUp/#`**:
  - **Description**: Topic used to send "wake-up" instructions to units in a low-power state.
  - **Rationale**: Ensures that units can be activated as needed to process events or collect data, while maintaining energy efficiency.

- **`/Data/From/#`**:
  - **Description**: The coordinator subscribes to this topic to receive data sent by the field units.
  - **Rationale**: Allows the coordinator to aggregate and process the received data from the units and forward it as needed.

- **`/Data/To/Unit/{ID}`**:
  - **Description**: Topic used to send specific data to a unit identified by its ID.
  - **Rationale**: Facilitates sending data or instructions to specific units, allowing granular control of operations.

### Coordinator Operation

1. **Event Monitoring**:
   - The coordinator subscribes to `/Control/Event/#` to monitor control events.
   - Upon receiving an event, it publishes corresponding instructions to `/Data/Coordinator/`.

2. **Receiving and Processing Data**:
   - The coordinator receives data on `/Data/Coordinator/#`.
   - It checks if the units are awake and, if necessary, sends wake-up instructions via `/Control/WakeUp/#`.
   - It forwards the data to the appropriate units via `/Data/To/Unit/{ID}`.

3. **Unit Wake-Up**:
   - The coordinator monitors `/Control/WakeUp/#` to send activation signals to the units.
   - Upon waking, the units may send data back through `/Data/From/#`.

---

## Communication Unit (ESP32)

### Overview

The communication units are ESP32 devices that connect to the coordinator via MQTT to receive instructions, collect data, and transmit it to the coordinator or external systems. They can enter a low-power state and wake up on demand as instructed by the coordinator.

### MQTT Topics

- **`/Data/To/Unit/{ID}`**:
  - **Description**: Topic where the unit subscribes to receive data or specific instructions from the coordinator.
  - **Rationale**: Allows the unit to receive data or commands directly from the coordinator, facilitating the execution of specific tasks.

- **`/Data/From/Unit/{ID}`**:
  - **Description**: Topic where the unit publishes processed data to be sent back to the coordinator.
  - **Rationale**: Facilitates the transmission of collected or processed data by the unit, ensuring that the coordinator can receive and process this information.

### Communication Unit Operation

1. **Connection and Subscription**:
   - The unit connects to the MQTT broker and subscribes to the `/Data/To/Unit/{ID}` topic to receive data and instructions.
   - It maintains a persistent session and uses QoS 1 to ensure data delivery.

2. **Receiving Data**:
   - Upon receiving a message on `/Data/To/Unit/{ID}`, the unit processes the message and retransmits the data to the external system via `RelayData()`.

3. **Sending Data to the Coordinator**:
   - After processing the data, the unit sends the results back to the coordinator by publishing to the `/Data/From/Unit/{ID}` topic.

4. **Maintaining Connection**:
   - The unit continuously checks the connection to the MQTT broker and reconnects automatically if disconnected.
   - It uses QoS 1 to ensure messages are delivered at least once, ensuring system reliability.

---

### For Collaborators

1. **Install `BlackBox`:**
   - **macOS**:
     ```bash
     brew install blackbox
     ```
   - **Linux (Debian/Ubuntu)**:
     ```bash
     sudo apt-get install blackbox
     ```
   - **Windows**:
     - On Windows, you can use `WSL` (Windows Subsystem for Linux) to install `BlackBox` the same way as on Linux, or set up a Unix-like terminal environment using tools like Git Bash and follow the Linux installation steps.

2. **Retrieve the Encrypted GPG Key:**
   - The GPG key is stored as a secret in the repository via AWS Secrets Manager.
   - Request access from the project maintainer to retrieve the GPG key, or follow the internal project guidelines to access it.

3. **Import the GPG Key on Your Machine:**
   - Once you have obtained the GPG key, import it using the following command:
     ```bash
     echo "base64_encoded_key" | base64 --decode | gpg --import
     ```
   - Replace `"base64_encoded_key"` with the actual base64-encoded GPG key value provided to you.

4. **Verify the Key Import:**
   - After importing, verify that the key was successfully added by running:
     ```bash
     gpg --list-keys
     ```
   - You should see the imported GPG key in the list of available keys.

5. **Decrypt Files Using `BlackBox`:**
   - With the GPG key imported, run the following command to decrypt all files in the repository:
     ```bash
     blackbox_decrypt_all_files
     ```
   - This will decrypt all protected files, making them accessible for your work.
