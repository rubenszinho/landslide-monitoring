# Communication System for Landslide Disaster Monitoring

## Project Objective

This project aims to implement a multiprotocol communication system for landslide disaster monitoring, using ESP32 devices as communication units and a central coordinator to manage communications. The system uses the MQTT protocol to ensure reliable and efficient communication between the units and the coordinator.

---

## Coordinator

### Overview

The coordinator is responsible for managing communication between the various field units (ESP32) and processing the received data. It receives control instructions, sends "wake-up" signals to the units as needed, and forwards processed data to the external system.

### MQTT Topics

- **`/Control/Event/#`**: 
  - **Description**: The coordinator subscribes to this topic to monitor control events that may require specific actions, such as instructions to wake up units or change the system state.
  - **Rationale**: This topic allows the coordinator to respond to global events that can affect the operation of the entire system.

- **`/Data/Coordinator/#`**:
  - **Description**: Topic used by the coordinator to receive data or specific instructions from units or other control systems.
  - **Rationale**: Centralizes the receipt of data and instructions for the coordinator, facilitating the processing and dissemination of information to the units.

- **`/Control/WakeUp/#`**:
  - **Description**: Topic used to send "wake-up" instructions to units that are in low-power mode.
  - **Rationale**: Ensures that the units can be activated as needed to process events or collect data, maintaining energy efficiency.

- **`/Data/From/#`**:
  - **Description**: The coordinator subscribes to this topic to receive data sent by the field units.
  - **Rationale**: Allows the coordinator to aggregate and process data received from the units, forwarding it as needed.

- **`/Data/To/Unit/{ID}`**:
  - **Description**: Topic used to send specific data to a unit identified by its ID.
  - **Rationale**: Facilitates sending data or instructions targeted at specific units, enabling granular control of operations.

### Coordinator Operation

1. **Event Monitoring**:
   - The coordinator subscribes to `/Control/Event/#` to monitor control events.
   - Upon receiving an event, it publishes corresponding instructions on `/Data/Coordinator/`.

2. **Data Reception and Processing**:
   - The coordinator receives data on `/Data/Coordinator/#`.
   - It checks if the units are awake; if necessary, it sends "wake-up" instructions on `/Control/WakeUp/#`.
   - Forwards the data to the appropriate units on `/Data/To/Unit/{ID}`.

3. **Unit Wake-Up**:
   - The coordinator monitors `/Control/WakeUp/#` to send activation signals to the units.
   - Upon waking up, the units can send data back via `/Data/From/#`.

---

## Communication Unit (ESP32)

### Overview

The communication units are ESP32 devices that connect to the coordinator via MQTT to receive instructions, collect data, and transmit it to the coordinator or external systems. They are capable of entering low-power states and waking up on demand, according to the coordinator's instructions.

### MQTT Topics

- **`/Data/To/Unit/{ID}`**:
  - **Description**: Topic where the unit subscribes to receive specific data or instructions from the coordinator.
  - **Rationale**: Allows the unit to receive data or commands directly from the coordinator, facilitating the execution of specific tasks.

- **`/Data/From/Unit/{ID}`**:
  - **Description**: Topic where the unit publishes processed data to be sent back to the coordinator.
  - **Rationale**: Facilitates the sending of data collected or processed by the unit, ensuring that the coordinator can receive and process this information.

### Communication Unit Operation

1. **Connection and Subscription**:
   - The unit connects to the MQTT broker and subscribes to the topic `/Data/To/Unit/{ID}` to receive data and instructions.
   - It maintains a persistent session and uses QoS 1 to ensure data delivery.

2. **Data Reception**:
   - Upon receiving a message on the topic `/Data/To/Unit/{ID}`, the unit processes the message and retransmits the data to the external system via `RelayData()`.

3. **Sending Data to the Coordinator**:
   - After processing the data, the unit sends the results back to the coordinator by publishing to the topic `/Data/From/Unit/{ID}`.

4. **Connection Maintenance**:
   - The unit continuously checks the connection to the MQTT broker and automatically reconnects in case of disconnection.
   - It uses QoS 1 to ensure that messages are delivered at least once, ensuring system reliability.

---

## For Collaborators

### Submodules

**Updating Submodules After Cloning**:
   - After cloning the repository, make sure to initialize and update any submodules by running the following command:
     ```bash
     git submodule update --init --recursive
     ```
   - This ensures that all submodules are correctly fetched and set up in your local repository.

### Blackbox

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

2. **Obtain the Encoded GPG Key:**
   - The base64-encoded GPG key is stored in the repository's "Secrets."
   - Ask the project maintainer to share the key with you if you do not have access yet.

3. **Import the GPG Key to Your Machine:**
   - After obtaining the key, import it using the following command:
     ```bash
     echo "base64_encoded_key" | base64 --decode | gpg --import
     ```
   - Replace `"base64_encoded_key"` with the base64 key you received.

4. **Verify the Import:**
   - You can verify if the key was successfully imported with the command:
     ```bash
     gpg --list-keys
     ```
   - The key should appear in the list of available keys.

5. **Decrypt Files with `BlackBox`:**
   - With the GPG key imported, use the following command to decrypt files:
     ```bash
     blackbox_decrypt_all_files
     ```
   - This will decrypt all protected files in the repository, allowing you to work with them.