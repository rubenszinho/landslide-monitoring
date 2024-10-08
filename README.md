# Communication System for Landslide Disaster Monitoring

## Project Objective

This project aims to implement a multiprotocol communication system for landslide disaster monitoring, using ESP32 devices as communication units and a central coordinator to manage communications. The system uses the MQTT protocol to ensure reliable and efficient communication between the units and the coordinator.

---

## Coordinator

### Overview

The coordinator is responsible for managing sensors directly connected to it (e.g., via GPIO pins on a Raspberry Pi). It collects data from the sensors, processes it, and publishes the data to MQTT topics. The coordinator also manages system control, such as waking up communication units, sending restart and shutdown signals, and monitoring its own health status.

### MQTT Topics

- **`/data/coordinator/sensor/{SensorType}/{SensorID}/{MeasurementType}`**:  
  - **Description**: The coordinator publishes sensor data to this topic. `{SensorType}` specifies the type of sensor (e.g., `soil`, `pluviometer`), `{SensorID}` uniquely identifies the sensor, and `{MeasurementType}` specifies the type of measurement (e.g., `temperature`, `humidity`).  
  - **Rationale**: Organizes sensor data based on its type, ID, and measurement, allowing external systems (e.g., communication units) to subscribe to the relevant data streams.

- **`/data/coordinator/health/{Metric}`**:  
  - **Description**: The coordinator publishes its own health status, including CPU, memory, and other operational metrics.  
  - **Rationale**: Allows external systems to monitor the coordinator's hardware health, ensuring proper operation and proactive maintenance.

- **`/control/wakeup/communication_unit/{CommUnitID}`**:  
  - **Description**: Topic used by the coordinator to send "wake-up" instructions to communication units.  
  - **Rationale**: Ensures that communication units can be activated from low-power mode when data needs to be transmitted or operations resumed.

- **`/control/restart/coordinator/{RestartType}`**:  
  - **Description**: Topic for issuing system restart commands to the coordinator. `{RestartType}` specifies the type of restart (e.g., `soft`, `full`).  
  - **Rationale**: Allows the coordinator to recover from potential deadlocks or software issues by performing a remote restart.

- **`/control/restart/communication_unit/{CommUnitID}/{RestartType}`**:  
  - **Description**: Topic for issuing restart commands to communication units. `{RestartType}` specifies the type of restart (e.g., `soft`, `full`).  
  - **Rationale**: Enables the system to remotely restart individual communication units for maintenance or error recovery.

- **`/control/shutdown/coordinator/{ShutdownType}`**:  
  - **Description**: Topic used to issue shutdown commands to the coordinator. `{ShutdownType}` specifies whether the shutdown should be `graceful` or `immediate`.  
  - **Rationale**: Facilitates controlled power management of the coordinator during periods of inactivity or maintenance.

- **`/control/shutdown/communication_unit/{CommUnitID}/{ShutdownType}`**:  
  - **Description**: Topic used to issue shutdown commands to communication units. `{ShutdownType}` specifies whether the shutdown should be `graceful` or `immediate`.  
  - **Rationale**: Facilitates controlled power management of communication units during periods of inactivity or maintenance.

### Coordinator Operation

1. **Data Publishing**:  
   The coordinator collects sensor data from connected devices (e.g., soil sensors, pluviometers) and publishes this data on `/data/coordinator/sensor/{SensorType}/{SensorID}/{MeasurementType}`. The communication units or other clients subscribe to these topics to receive the sensor data.

2. **Health Monitoring**:  
   The coordinator publishes its hardware health metrics to `/data/coordinator/health/{Metric}` for external monitoring. Metrics may include CPU usage, memory availability, and battery life (if applicable).

3. **Control Operations**:  
   The coordinator sends wake-up signals to communication units on `/control/wakeup/communication_unit/{CommUnitID}` and manages system-wide restarts and shutdowns through the `/control/restart/#` and `/control/shutdown/#` topics. The coordinator itself can also be restarted or shut down when necessary.

---

### Submodules

**Updating Submodules After Cloning**:
   - After cloning the repository, make sure to initialize and update any submodules by running the following command:
     ```bash
     git submodule update --init --recursive
     ```
   - This ensures that all submodules are correctly fetched and set up in your local repository.

   - To keep your submodules up-to-date with their branches during your workflow, you could:
     ```bash
     git submodule update --remote
     ```

### Setting up Blackbox for File Encryption

To securely encrypt sensitive files in this project using **StackExchange Blackbox**, follow these steps:

#### 1. Install `BlackBox`
You can automatically install StackExchange Blackbox via the following commands:

```bash
git clone https://github.com/StackExchange/blackbox.git
cd blackbox
sudo make copy-install
```
This will copy the necessary files into `/usr/local/bin`.

#### 2. Obtain the Encoded GPG Keys
The **public** and **private** Base64-encoded GPG keys are stored in the repository's "Secrets." 
Ask the project maintainer to share the keys with you if you do not have access yet.

You will receive:
- A **Base64-encoded public key**
- A **Base64-encoded private key**

#### 3. Import the Public Key
Once you receive the **Base64-encoded public key**, use the following command to decode and import it:

```bash
echo "base64_encoded_public_key" | base64 --decode | gpg --import
```

- Replace `base64_encoded_public_key` with the actual Base64-encoded string of the public key.

#### 4. Import the Private Key
After importing the public key, you'll also need to import the **private key** for decryption purposes. To do that, use the following command:

```bash
echo "base64_encoded_private_key" | base64 --decode | gpg --import
```

- Replace `base64_encoded_private_key` with the actual Base64-encoded string of the private key.

#### 5. Verify the Import
You can verify if both keys were successfully imported with the following command:

```bash
gpg --list-secret-keys
```

This will list the GPG keys on your system, and you should see both the public and private key associated with your GPG email.

#### 6. Decrypt Files with `BlackBox`
With both the public and private keys imported, you can now decrypt the files in your project:

```bash
blackbox_decrypt_all_files
```

This command will decrypt all files that were encrypted with Blackbox, using your imported GPG keys.
