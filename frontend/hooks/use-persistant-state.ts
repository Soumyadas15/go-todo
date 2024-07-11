import { useState, useEffect, useRef } from 'react';

export const usePersistantState = (name: string, defaultValue: any) => {
  const [value, setValue] = useState(defaultValue);
  const nameRef = useRef(name);

  useEffect(() => {
    try {
      const storedValue = localStorage.getItem(name);
      if (storedValue !== null) setValue(JSON.parse(storedValue));
      else localStorage.setItem(name, JSON.stringify(defaultValue));
    } catch (error) {
      console.error('Error loading or setting localStorage:', error);
      setValue(defaultValue);
    }
  }, [name, defaultValue]);

  useEffect(() => {
    try {
      localStorage.setItem(nameRef.current, JSON.stringify(value));
    } catch (error) {
      console.error('Error saving to localStorage:', error);
    }
  }, [value]);

  useEffect(() => {
    const lastName = nameRef.current;
    if (name !== lastName) {
      try {
        localStorage.setItem(name, JSON.stringify(value));
        nameRef.current = name;
        localStorage.removeItem(lastName);
      } catch (error) {
        console.error('Error renaming or removing from localStorage:', error);
      }
    }
  }, [name, value]);

  return [value, setValue];
};
