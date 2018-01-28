import React, { Component } from 'react';
import {
  AppRegistry,
  StyleSheet,
  Text,
  View,
  Image,
  TouchableOpacity,
  TouchableHighlight,
  ScrollView,
} from 'react-native';

import ModalDropdown from 'react-native-modal-dropdown';

const DEMO_OPTIONS_1 = ['option 1', 'option 2'];
const DEMO_OPTIONS_2 = [
  {"name": "Rex"},
  {"name": "Mary"}
];
const DEMO_OPTIONS_3 = [
  {"name": "John"},
  {"name": "Steve"}
];

export default class MyEvents extends Component {
  constructor(props) {
    super(props);

    this.state = {
      dropdown_4_options: null,
      dropdown_4_defaultValue: 'loading...',
      dropdown_6_icon_heart: true,
    };
  }

  render() {
    return (
      <View style={styles.container}>
      	<View style={styles.row}>
            <ModalDropdown style={styles.dropdown_2}
                           textStyle={styles.dropdown_2_text}
                           dropdownStyle={styles.dropdown_2_dropdown}
                           options={DEMO_OPTIONS_2}
                           renderButtonText={(rowData) => this._dropdown_2_renderButtonText(rowData)}
                           renderRow={this._dropdown_2_renderRow.bind(this)}
                           renderSeparator={(sectionID, rowID, adjacentRowHighlighted) => this._dropdown_2_renderSeparator(sectionID, rowID, adjacentRowHighlighted)}
            ><Text style={{fontSize:30, color: 'white', textAlign: 'center'}}>Concert</Text></ModalDropdown>
        </View>
        <View style={styles.row}>
            <ModalDropdown style={styles.dropdown_2}
                           textStyle={styles.dropdown_2_text}
                           dropdownStyle={styles.dropdown_2_dropdown}
                           options={DEMO_OPTIONS_3}
                           renderButtonText={(rowData) => this._dropdown_2_renderButtonText(rowData)}
                           renderRow={this._dropdown_2_renderRow.bind(this)}
                           renderSeparator={(sectionID, rowID, adjacentRowHighlighted) => this._dropdown_2_renderSeparator(sectionID, rowID, adjacentRowHighlighted)}
            ><Text style={{fontSize:30, color: 'white', textAlign: 'center'}}>Star Wars</Text></ModalDropdown>
        </View>
      </View>
    );
  }

  _dropdown_2_renderButtonText(rowData) {
    const { name } = rowData;
    return `${name}`;
  }

  _dropdown_2_renderRow(rowData, rowID, highlighted) {
    let icon = require('./user.png');
    let evenRow = rowID % 2;
    return (
      <TouchableHighlight underlayColor='cornflowerblue'>
        <View style={[styles.dropdown_2_row, {backgroundColor: evenRow ? 'lemonchiffon' : 'white'}]}>
          <Image style={styles.dropdown_2_image}
                 mode='stretch'
                 source={icon}
          />
          <Text style={[styles.dropdown_2_row_text, highlighted && {color: 'mediumaquamarine'}]}>
            {`${rowData.name}`}
          </Text>
        </View>
      </TouchableHighlight>
    );
  }

  _dropdown_2_renderSeparator(sectionID, rowID, adjacentRowHighlighted) {
    if (rowID == DEMO_OPTIONS_1.length - 1) return;
    let key = `spr_${rowID}`;
    return (<View style={styles.dropdown_2_separator}
                  key={key}
    />);
  }

 
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  row: {
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center'
  },
  cell: {
    flex: 1,

  },
  scrollView: {
    flex: 1,
  },
  contentContainer: {
    height: 500,
    paddingVertical: 100,
    paddingLeft: 20,
  },
  textButton: {
    color: 'deepskyblue',
    borderColor: 'deepskyblue',
    margin: 2,
  },

  dropdown_2: {
    width: 150,
    borderWidth: 0,
    borderRadius: 3,
    backgroundColor: 'cornflowerblue',
  },
  dropdown_2_text: {
    marginVertical: 10,
    marginHorizontal: 6,
    fontSize: 18,
    color: 'white',
    textAlign: 'center',
    textAlignVertical: 'center',
  },
  dropdown_2_dropdown: {
    width: 150,
    height: 90,
    borderColor: 'cornflowerblue',
    borderWidth: 2,
    borderRadius: 3,
  },
  dropdown_2_row: {
    flexDirection: 'row',
    height: 40,
    alignItems: 'center',
  },
  dropdown_2_image: {
    marginLeft: 4,
    width: 30,
    height: 30,
  },
  dropdown_2_row_text: {
    marginHorizontal: 4,
    fontSize: 16,
    color: 'navy',
    textAlignVertical: 'center',
  },
  dropdown_2_separator: {
    height: 1,
    backgroundColor: 'cornflowerblue',
  }
});