import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import styles from './styles.js'

export default class App extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.navbar}>
          <Text style={{fontFamily: 'Cochin', fontSize: 20, marginTop: 25}}>Select Event</Text>
        </View>
      </View>

    );
  }
}
